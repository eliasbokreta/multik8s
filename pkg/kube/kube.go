package kube

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

const (
	kubeconfigDir      = ".kube"
	kubeconfigFilename = "config"
)

// Create a Configclient
func createConfigClient(kubeconfigPath string) clientcmd.ClientConfig {
	if kubeconfigPath == "" {
		kubeconfigPath = filepath.Join(homedir.HomeDir(), kubeconfigDir, kubeconfigFilename)
	}
	return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{},
	)
}

// Create a Rawconfig
func GetRawAPIConfig(kubeconfigPath string) (*api.Config, error) {
	configclient := createConfigClient(kubeconfigPath)

	rawConfig, err := configclient.RawConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to get kube rawconfig: %w", err)
	}

	return &rawConfig, nil
}

// Create a new Clientset
func GetClientSet(kubeconfigPath string) (*kubernetes.Clientset, error) {
	configclient := createConfigClient(kubeconfigPath)

	clientcfg, err := configclient.ClientConfig()
	if err != nil {
		return nil, fmt.Errorf("could not create clientconfig: %w", err)
	}

	clientset, err := kubernetes.NewForConfig(clientcfg)
	if err != nil {
		return nil, fmt.Errorf("could not create clientset: %w", err)
	}

	return clientset, nil
}
