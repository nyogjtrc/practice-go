package main

import (
	"fmt"
	"os"

	"github.com/nyogjtrc/practice-go/raking_with_redis/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
