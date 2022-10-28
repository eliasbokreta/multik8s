package kube

import (
	"bufio"
	"context"
	"strings"
	"sync"
	"time"

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
	Follow      bool
	TailLines   int64
}

type PodInfo struct {
	ContextID       int
	Cluster         string
	Namespace       string
	Podname         string
	PodGenerateName string
	PodID           int
	Phase           string
	Age             string
}

// List all pods on a given namespace
func PodList(cfg Config, wg *sync.WaitGroup, podlist *[]PodInfo) {
	defer wg.Done()

	tmpKubeconfigPath, err := GenerateTemporaryKubeconfig(cfg.KubeContext)
	if err != nil {
		log.Errorf("could not generate temporary kubeconfig: %v", err)
		return
	}

	clientset, err := GetClientSet(tmpKubeconfigPath)
	if err != nil {
		log.Errorf("could not get clientset: %v", err)
		return
	}

	apiconfig, err := GetRawAPIConfig(tmpKubeconfigPath)
	if err != nil {
		log.Errorf("could not get raw API config: %v", err)
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

		seconds := uint64(time.Since(pod.Status.StartTime.Time).Seconds())
		*podlist = append(*podlist, PodInfo{
			Cluster:         apiconfig.Contexts[apiconfig.CurrentContext].Cluster,
			Namespace:       pod.Namespace,
			Podname:         pod.Name,
			PodGenerateName: pod.GenerateName,
			Phase:           string(pod.Status.Phase),
			Age:             utils.AgeFormatter(seconds),
			ContextID:       cfg.ID,
		})
	}
}

// Read pod logs
// Credits: https://github.com/nwaizer/GetPodLogsEfficiently
//
//nolint:cyclop,funlen
func PodLogs(cancelCtx context.Context, cfg Config, wg *sync.WaitGroup) {
	defer wg.Done()

	tmpKubeconfigPath, err := GenerateTemporaryKubeconfig(cfg.KubeContext)
	if err != nil {
		log.Errorf("could not generate temporary kubeconfig: %v", err)
		return
	}

	clientset, err := GetClientSet(tmpKubeconfigPath)
	if err != nil {
		log.Errorf("could not get clientset: %v", err)
		return
	}

	apiconfig, err := GetRawAPIConfig(tmpKubeconfigPath)
	if err != nil {
		log.Errorf("could not get raw API config: %v", err)
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
		podLogOptions := corev1.PodLogOptions{
			Follow:     cfg.Follow,
			TailLines:  &cfg.TailLines,
			Timestamps: true,
		}

		if !cfg.Follow {
			podLogs := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOptions)

			stream, err := podLogs.Stream(context.Background())
			if err != nil {
				return
			}
			defer stream.Close()

			reader := bufio.NewScanner(stream)
			for reader.Scan() {
				outputLogs(cfg.ID, apiconfig.Contexts[apiconfig.CurrentContext].Cluster, pID, pod, reader.Text())
			}
			continue
		}

		w.Add(1)
		go func(p corev1.Pod, podID int) {
			defer w.Done()
			podLogs := clientset.CoreV1().Pods(p.Namespace).GetLogs(p.Name, &podLogOptions)

			stream, err := podLogs.Stream(context.Background())
			if err != nil {
				return
			}
			defer stream.Close()

			reader := bufio.NewScanner(stream)
			for reader.Scan() {
				outputLogs(cfg.ID, apiconfig.Contexts[apiconfig.CurrentContext].Cluster, podID, p, reader.Text())
			}
		}(pod, pID)
	}
	w.Wait()
}
