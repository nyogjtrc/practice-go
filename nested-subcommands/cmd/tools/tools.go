package tools

import (
	"github.com/spf13/cobra"
)

// ToolsCmd is a collection of tool commands
var ToolsCmd = &cobra.Command{
	Use:   "tools",
	Short: "a collection of tool commands",
	Long:  `a collection of tool commands`,
}

func init() {
}
