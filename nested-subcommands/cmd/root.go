package cmd

import (
	"os"

	"github.com/nyogjtrc/practice-go/nested-subcommands/cmd/tools"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A nested subcommands app",
	Long:  `A nested subcommands app`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(tools.ToolsCmd)
}
