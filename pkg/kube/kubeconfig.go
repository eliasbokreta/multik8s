package kube

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/util/homedir"
)

const (
	kubeConfigDir      = ".kube"
	kubeConfigFilename = "config"
)

// Get Rawconfig from kubeconfig file
func GetRawConfig() (*api.Config, error) {
	kubeConfigPath := filepath.Join(homedir.HomeDir(), kubeConfigDir, kubeConfigFilename)
	config := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfigPath},
		&clientcmd.ConfigOverrides{},
	)

	rawConfig, err := config.RawConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to get kube rawconfig: %w", err)
	}

	return &rawConfig, nil
}

// Generate a temporary kubeconfig file
func GenerateTemporaryKubeconfig(context string) (string, error) {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "multik8s")
	if err != nil {
		return "", fmt.Errorf("could not create tempfile")
	}

	filePath := fmt.Sprintf("%s/config.yaml", tmpDir)

	kubeConfigPath := filepath.Join(homedir.HomeDir(), kubeConfigDir, kubeConfigFilename)
	kubeconfig, err := clientcmd.LoadFromFile(kubeConfigPath)
	if err != nil {
		return "", fmt.Errorf("could not load from file: %w", err)
	}
	kubeconfig.CurrentContext = context

	if err := clientcmd.WriteToFile(*kubeconfig, filePath); err != nil {
		return "", fmt.Errorf("could not write to file: %w", err)
	}

	return filePath, nil
}
