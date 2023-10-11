package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	for {

		topic := client.Topic(topicName)
		res := topic.Publish(ctx, &pubsub.Message{
			Data: []byte("hello world" + time.Now().Format(time.RFC3339)),
		})
		// The publish happens asynchronously.
		// Later, you can get the result from res:
		msgID, err := res.Get(ctx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("msgID", msgID)
		time.Sleep(100 * time.Millisecond)
	}
}
