package cmd

import (
	"github.com/labstack/echo"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "serv",
	Short: "s",
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
}

func server() {
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
