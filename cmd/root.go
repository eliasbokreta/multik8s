package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "multik8s",
	Short: "multik8s is an utility Kubernetes",
	Long:  "multik8s is an utility tool that allow fetching data from multiple kube context at once"}

func initCmd() {
	cobra.OnInitialize()
	multik8sInit()
}

func Execute() error {
	initCmd()
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("could not run the command tree: %w", err)
	}

	return nil
}
