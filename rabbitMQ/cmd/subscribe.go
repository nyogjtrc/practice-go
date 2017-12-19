package cmd

import (
	"log"

	"github.com/nyogjtrc/practice-go/rabbitMQ/hunter"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "subscribe",
	Run: func(cmd *cobra.Command, args []string) {
		sub()
	},
}

func init() {
	RootCmd.AddCommand(subCmd)
}

func sub() {
	h := hunter.New("amqp://guest:guest@localhost:5672/", "hello")
	err := h.Connect()
	failOnError(err, "Failed to connect to RabbitMQ")
	defer h.Close()

	msgs, err := h.Consume()

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
