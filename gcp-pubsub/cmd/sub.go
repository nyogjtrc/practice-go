package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/pubsub"
)

const projectID = ""

const topicName = ""

func main() {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	// Use a callback to receive messages via subscription1.
	sub := client.Subscription("nyo-test-topic-sub")
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println(string(m.Data))
		m.Ack() // Acknowledge that we've consumed the message.
	})
	if err != nil {
		log.Println(err)
	}
}
