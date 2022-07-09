package multik8s

import (
	"context"
	"fmt"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"github.com/eliasbokreta/multik8s/pkg/kube"
	"github.com/eliasbokreta/multik8s/pkg/utils"
	log "github.com/sirupsen/logrus"
	_ "k8s.io/client-go/plugin/pkg/client/auth/azure" // Import needed for Azure authentication if using AKS
)

type Config struct {
	SelectedContexts []string
	Namespace        string
	PodName          string
	Follow           bool
	TailLines        int64
}

// Create a new MultiK8s config
func New(config Config) *Config {
	return &Config{
		Namespace: config.Namespace,
		PodName:   config.PodName,
		Follow:    config.Follow,
		TailLines: config.TailLines,
	}
}

// Generate a select menu to choose Kubernetes contexts
func (c *Config) SelectContexts() error {
	promptValues := []string{}

	rawConfig, err := kube.GetRawConfig()
	if err != nil {
		return fmt.Errorf("could not get raw config: %w", err)
	}

	for cluster := range rawConfig.Contexts {
		promptValues = append(promptValues, cluster)
	}

	prompt := &survey.MultiSelect{
		Message: "Which context(s) do you want to select ?",
		Options: promptValues,
	}

	err = survey.AskOne(prompt, &c.SelectedContexts, survey.WithKeepFilter(true))
	if err != nil {
		return fmt.Errorf("cannot generate survey: %w", err)
	}

	return nil
}

// MultiK8s entrypoint
func (c *Config) Run(action string) error {
	if err := c.SelectContexts(); err != nil {
		return fmt.Errorf("could not select context: %w", err)
	}
	log.Infof("Starting fetching data for contexts: %v", c.SelectedContexts)

	var wg sync.WaitGroup
	ctx := context.Background()
	cancelCtx, endGoFunc := context.WithCancel(ctx)

	podList := make([]kube.PodInfo, 0)

	for id, ctx := range c.SelectedContexts {
		wg.Add(1)

		cfg := kube.Config{
			ID:          id,
			KubeContext: ctx,
			Namespace:   c.Namespace,
			PodName:     c.PodName,
			Follow:      c.Follow,
			TailLines:   c.TailLines,
		}

		switch action {
		case "podLogs":
			go kube.PodLogs(cancelCtx, cfg, &wg)
		case "podList":
			go kube.PodList(cfg, &wg, &podList)
		}
	}
	wg.Wait()
	endGoFunc()

	if action == "podList" {
		table := utils.GetTableWriter([]string{"Context", "Namespace", "Pod", "Status"})

		for _, pod := range podList {
			line := []string{pod.Cluster, pod.Namespace, pod.Podname, pod.Phase}
			table.Append(line)
		}
		table.Render()
	}

	return nil
}
