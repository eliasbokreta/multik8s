package kube

import (
	"fmt"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

// Generate Kubernetes clientset and clientrest
func GetKubernetesClients(ctx string) (*kubernetes.Clientset, *api.Config, error) {
	tempKubeconfig, err := GenerateTemporaryKubeconfig(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("could not generate temporary kubeconfig: %w", err)
	}
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: tempKubeconfig},
		&clientcmd.ConfigOverrides{},
	)

	clientcfg, err := config.ClientConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("could not create clientconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(clientcfg)
	if err != nil {
		return nil, nil, fmt.Errorf("could not create clientset: %w", err)
	}

	clientrest, err := config.RawConfig()
	if err != nil {
		return nil, nil, fmt.Errorf("could not create clientrest: %w", err)
	}

	return clientset, &clientrest, nil
}
