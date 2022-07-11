package cmd

import (
	"fmt"
	"os"

	"github.com/eliasbokreta/multik8s/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var version string

var rootCmd = &cobra.Command{
	Use:     "multik8s",
	Short:   "multik8s is a Kubernetes utility CLI tool",
	Long:    "multik8s is an utility CLI tool that allow fetching data from multiple Kubernetes context at once",
	Version: version,
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update multik8s version",
	Long:  "update multik8s to latest version",
	Run: func(cmd *cobra.Command, args []string) {
		if err := utils.SelfUpdate(version); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var docCmd = &cobra.Command{
	Use:    "doc",
	Short:  "multik8s cmd documentation",
	Long:   "multik8s commands' markdown documentation",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		if err := doc.GenMarkdownTree(rootCmd, "./docs"); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func initCmd() {
	cobra.OnInitialize()
	multik8sCmdInit()
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(docCmd)
}

func Execute() error {
	initCmd()
	if err := rootCmd.Execute(); err != nil {
		return fmt.Errorf("could not run the command tree: %w", err)
	}

	return nil
}
