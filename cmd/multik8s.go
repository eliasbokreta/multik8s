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
	Short: "Get <pods | logs> subcommand",
}

var cmdGetPods = &cobra.Command{
	Use:   "pods",
	Short: "get pods",
	Long:  "get pods basic information",
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
	Short: "get logs",
	Long:  "get pods logs",
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

func multik8sCmdInit() {
	rootCmd.AddCommand(cmdGet)
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "Kubernetes namespace (should be the exact name)")
	rootCmd.PersistentFlags().StringVarP(&podName, "podname", "p", "", "Kubernetes pod name (works as a wildcard)")
	cmdGet.AddCommand(cmdGetPods)
	cmdGetLogs.PersistentFlags().BoolVarP(&follow, "follow", "f", false, "Choose whether or not to follow log stream")
	cmdGetLogs.PersistentFlags().Int64VarP(&tailLines, "tail", "t", 5, "The number of lines from the end of the logs to show")
	cmdGet.AddCommand(cmdGetLogs)
}
