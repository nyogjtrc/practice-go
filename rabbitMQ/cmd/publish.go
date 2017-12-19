package cmd

import (
	"fmt"
	"log"

	"github.com/nyogjtrc/practice-go/rabbitMQ/hunter"

	"github.com/spf13/cobra"
)

var publishCmd = &cobra.Command{
	Use:   "pub",
	Short: "publish",
	Run: func(cmd *cobra.Command, args []string) {
		pub()
	},
}

func init() {
	RootCmd.AddCommand(publishCmd)
}

func pub() {
	h := hunter.New("amqp://guest:guest@localhost:5672/", "hello")
	err := h.Connect()
	failOnError(err, "Failed to connect to RabbitMQ")
	defer h.Close()

	body := []byte("this is a push message")
	err = h.Publish(body)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
