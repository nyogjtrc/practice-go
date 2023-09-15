package tools

import (
	"fmt"

	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "tool: monitor command",
	Long:  `tool: monitor command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("monitor called")
	},
}

func init() {
	ToolsCmd.AddCommand(monitorCmd)
}
