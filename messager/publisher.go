package messager

import (
	"os"

	"github.com/streadway/amqp"
)

func Publish(message amqp.Publishing) {
    
    amqpServerURL := os.Getenv("AMQP_SERVER_URL")
    
    // Create a new RabbitMQ connection.
    connectRabbitMQ, err := amqp.Dial(amqpServerURL)
    if err != nil {
        panic(err)
    }
    defer connectRabbitMQ.Close()

    // Let's start by opening a channel over the connection
    channelRabbitMQ, err := connectRabbitMQ.Channel()
    if err != nil {
        panic(err)
    }
    defer channelRabbitMQ.Close()

    // With the instance and declare Queues that we can
    // publish and subscribe to.
    _, err = channelRabbitMQ.QueueDeclare(
        "QueueService", // queue name
        true,            // durable
        false,           // auto delete
        false,           // exclusive
        false,           // no wait
        nil,             // arguments
    )

    if err != nil {
        panic(err)
    }

	// Publish a message to the queue.
	channelRabbitMQ.Publish(
		"",                 // exchange
		"QueueService",     // queue name
		false,              // mandatory
		false,              // immediate
		message,            // message to publish
	)
}