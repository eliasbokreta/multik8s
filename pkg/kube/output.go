package kube

import (
	"fmt"

	"github.com/eliasbokreta/multik8s/pkg/utils"
	"github.com/olekukonko/tablewriter"
	corev1 "k8s.io/api/core/v1"
)

// Print streamed logs
func outputLogs(contextID int, cluster string, podID int, pod corev1.Pod, line string) {
	contextColor, podColor := utils.GetColorsFn(contextID)
	podIDColor := utils.GetColorFn(podID)

	var podIDx string
	for i := len(pod.GenerateName); i < len(pod.Name); i++ {
		podIDx += string(pod.Name[i])
	}

	message := fmt.Sprintf("[%s][%s%s][%s]%s", contextColor(cluster), podColor(pod.GenerateName), podIDColor(podIDx), podIDColor(podID), line)
	fmt.Println(message)
}

// Print pod list as table
func OutputPods(podList []PodInfo, header []string) {
	table := utils.GetTableWriter(header)

	podNameUnique := make(map[string]bool)
	for _, pod := range podList {
		podNameUnique[pod.PodGenerateName] = true
	}

	pColors := make(map[string]int)
	count := 0
	for pod := range podNameUnique {
		pColors[pod] = count
		count++
	}

	for i, pod := range podList {
		if _, ok := pColors[pod.PodGenerateName]; ok {
			podList[i].PodID = pColors[pod.PodGenerateName]
		}
	}

	for _, pod := range podList {
		line := []string{
			pod.Cluster,
			pod.Namespace,
			pod.Podname,
			pod.Phase,
			pod.Age,
		}

		colors := []tablewriter.Colors{
			utils.GetTableRowColor(pod.ContextID),
			tablewriter.Color(),
			utils.GetTableRowColor(pod.PodID),
			tablewriter.Color(),
			tablewriter.Color(),
		}

		table.Rich(line, colors)
	}
	table.Render()
}
