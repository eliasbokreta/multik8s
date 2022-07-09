package cmd

import (
	"fmt"
	"os"

	"github.com/eliasbokreta/multik8s/pkg/multik8s"
	"github.com/spf13/cobra"
)

var (
	namespace string
	podName   string
	follow    bool
	tailLines int64
)

var cmdGet = &cobra.Command{
	Use:   "get",
	Short: "Kubernetes related commands",
}

var cmdGetPods = &cobra.Command{
	Use:   "pods",
	Short: "get pods",
	Long:  "Get pods",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := multik8s.Config{
			Namespace: namespace,
			PodName:   podName,
		}
		m := multik8s.New(cfg)
		if err := m.Run("podList"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var cmdGetLogs = &cobra.Command{
	Use:   "logs",
	Short: "get pod logs",
	Long:  "Get pod logs",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := multik8s.Config{
			Namespace: namespace,
			PodName:   podName,
			Follow:    follow,
			TailLines: tailLines,
		}
		m := multik8s.New(cfg)
		if err := m.Run("podLogs"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func multik8sInit() {
	rootCmd.AddCommand(cmdGet)
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "Kubernetes namespace")
	rootCmd.PersistentFlags().StringVarP(&podName, "podname", "p", "", "Kubernetes pod name")
	cmdGet.AddCommand(cmdGetPods)
	cmdGetLogs.PersistentFlags().BoolVarP(&follow, "follow", "f", false, "Follow log stream")
	cmdGetLogs.PersistentFlags().Int64VarP(&tailLines, "tail", "t", 5, "The number of lines from the end of the logs to show")
	cmdGet.AddCommand(cmdGetLogs)
}
