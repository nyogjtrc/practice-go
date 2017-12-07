package hunter

import (
	"github.com/streadway/amqp"
)

// RabbitHunter store data for connect rabbitMQ
type RabbitHunter struct {
	rabbitURL   string
	rConnection *amqp.Connection
	rChannel    *amqp.Channel
	rQueue      amqp.Queue
	QueueName   string
}

// New RabbitHunter
func New(url string, name string) *RabbitHunter {
	h := &RabbitHunter{
		rabbitURL: url,
		QueueName: name,
	}
	return h
}

// Close connection and channel
func (h *RabbitHunter) Close() {
	if h.rConnection != nil {
		defer h.rConnection.Close()
	}
	if h.rChannel != nil {
		h.rChannel.Close()
	}
}

// Connect dial to rabbitMQ, connect to channel, and declare queue
func (h *RabbitHunter) Connect() (err error) {
	h.rConnection, err = amqp.Dial(h.rabbitURL)
	if err != nil {
		return
	}

	h.rChannel, err = h.rConnection.Channel()
	if err != nil {
		return
	}

	return h.qqDeclare()
}

// qqDeclare declare a queue
func (h *RabbitHunter) qqDeclare() (err error) {
	h.rQueue, err = h.rChannel.QueueDeclare(
		h.QueueName, // name
		true,        // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	return
}

// Consume create rabbitMQ consumer
func (h *RabbitHunter) Consume() (<-chan amqp.Delivery, error) {
	var err error
	if err = h.Connect(); err != nil {
		return nil, err
	}

	err = h.rChannel.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return nil, err
	}

	return h.rChannel.Consume(
		h.QueueName, // queue
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
}

// Publish json []byte to rabbitMQ
func (h *RabbitHunter) Publish(body []byte) error {
	return h.rChannel.Publish(
		"",          // exchange
		h.QueueName, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		},
	)
}
