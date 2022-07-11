package kube

import (
	"fmt"
	"io/ioutil"
	"os"

	"k8s.io/client-go/tools/clientcmd"
)

// Generate a temporary kubeconfig file
func GenerateTemporaryKubeconfig(context string) (string, error) {
	tmpDir, err := ioutil.TempDir(os.TempDir(), "multik8s")
	if err != nil {
		return "", fmt.Errorf("could not create tempfile")
	}

	filePath := fmt.Sprintf("%s/config.yaml", tmpDir)

	rawAPIConfig, err := GetRawAPIConfig("")
	if err != nil {
		return "", fmt.Errorf("could not get raw config: %w", err)
	}
	rawAPIConfig.CurrentContext = context

	if err := clientcmd.WriteToFile(*rawAPIConfig, filePath); err != nil {
		return "", fmt.Errorf("could not write to file: %w", err)
	}

	return filePath, nil
}
