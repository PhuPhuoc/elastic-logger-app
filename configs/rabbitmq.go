package configs

import (
	"log"

	"github.com/streadway/amqp"
)

func ConnectRabbitMQ(config *Config) *amqp.Connection {
	conn, err := amqp.Dial(config.RABBITMQ_URL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	return conn
}

func DeclareExchange(ch *amqp.Channel, name string) {
	err := ch.ExchangeDeclare(
		name,
		"topic", // hoặc "direct", "fanout", v.v.
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}
}

func DeclareQueue(ch *amqp.Channel, name string) amqp.Queue {
	q, err := ch.QueueDeclare(
		name,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}
	return q
}
