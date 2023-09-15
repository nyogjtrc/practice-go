package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "web command",
	Long:  `web command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("web called")
	},
}

func init() {
	rootCmd.AddCommand(webCmd)
}
