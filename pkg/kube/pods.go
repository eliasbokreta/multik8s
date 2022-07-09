package kube

import (
	"bufio"
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/eliasbokreta/multik8s/pkg/utils"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Config struct {
	ID          int
	KubeContext string
	Namespace   string
	PodName     string
}

type PodInfo struct {
	Cluster   string
	Namespace string
	Podname   string
	Phase     string
}

// List all pods on a given namespace
func PodList(cfg Config, wg *sync.WaitGroup, podlist *[]PodInfo) {
	defer wg.Done()

	clientset, clientrest, err := GetKubernetesClients(cfg.KubeContext)
	if err != nil {
		log.Errorf("could not get kubernetes clients: %v", err)
		return
	}

	pods, err := clientset.CoreV1().Pods(cfg.Namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Errorf("could not list pods: %v", err)
		return
	}

	for _, pod := range pods.Items {
		if cfg.PodName != "" {
			if !strings.Contains(strings.ToLower(pod.Name), strings.ToLower(cfg.PodName)) {
				continue
			}
		}

		*podlist = append(*podlist, PodInfo{
			Cluster:   clientrest.Contexts[clientrest.CurrentContext].Cluster,
			Namespace: pod.Namespace,
			Podname:   pod.Name,
			Phase:     string(pod.Status.Phase),
		})
	}
}

// Read pod logs
// nolint: cyclop
func PodLogs(cancelCtx context.Context, cfg Config, wg *sync.WaitGroup) {
	defer wg.Done()

	clientset, clientrest, err := GetKubernetesClients(cfg.KubeContext)
	if err != nil {
		log.Errorf("could not get kubernetes clients: %v", err)
		return
	}

	pods, err := clientset.CoreV1().Pods(cfg.Namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Errorf("could not list pods: %v", err)
		return
	}

	var w sync.WaitGroup

	for pID, pod := range pods.Items {
		if cfg.PodName != "" {
			if !strings.Contains(strings.ToLower(pod.Name), strings.ToLower(cfg.PodName)) {
				continue
			}
		}

		w.Add(1)
		go func(p corev1.Pod, podID int) {
			defer w.Done()
			count := int64(100)
			podLogOptions := corev1.PodLogOptions{
				Follow:     true,
				TailLines:  &count,
				Timestamps: true,
			}
			podLogRequest := clientset.CoreV1().Pods(p.Namespace).GetLogs(p.Name, &podLogOptions)

			stream, err := podLogRequest.Stream(context.Background())
			if err != nil {
				return
			}
			defer stream.Close()

			reader := bufio.NewScanner(stream)
			line := ""
			for {
				for reader.Scan() {
					select {
					case <-cancelCtx.Done():
						break
					default:
						line = reader.Text()
						contextColor, podColor := utils.GetColorsFn(cfg.ID)
						podIDColor := utils.GetColorFn(podID)
						message := fmt.Sprintf("[%s][%s][%s]%s", contextColor(clientrest.Contexts[clientrest.CurrentContext].Cluster), podColor(p.Name), podIDColor(podID), line)
						fmt.Println(message)
					}
				}
			}
		}(pod, pID)
	}
	w.Wait()
}
