package consumer

import (
	"log"
	"os"
	"time"
	"vts_api/service"

	"github.com/streadway/amqp"
)

func ConsumerInit () {
    // Define RabbitMQ server URL.
    amqpServerURL := os.Getenv("AMQP_SERVER_URL")

    // Create a new RabbitMQ connection.
    connectRabbitMQ, err := amqp.Dial(amqpServerURL)
    if err != nil {
        panic(err)
    }
    defer connectRabbitMQ.Close()

    // Opening a channel over the connection already established.
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

    // Subscribing to QueueService1 for getting messages.
    messages, err := channelRabbitMQ.Consume(
        "QueueService", // queue name
        "",              // consumer
        true,            // auto-ack
        false,           // exclusive
        false,           // no local
        false,           // no wait
        nil,             // arguments
    )
    if err != nil {
        log.Println(err)
    }

    // Build a welcome message.
    log.Println("Successfully connected to RabbitMQ")
    log.Println("Waiting for messages")

    // Make a channel to receive messages into infinite loop.
    forever := make(chan bool)

    go func() {
        for message := range messages {
            log.Printf("Received message")
            log.Printf(" > %s\n", message.Body)

            if (!service.Notify(message.Body)) {
                
                time.Sleep(5 * time.Second)
                log.Printf("Retring after 5 seconds")
                
                if (!service.Notify(message.Body)) {
                    // Last try
                    time.Sleep(15 * time.Second)
                    log.Printf("Retring after 15 seconds")
                    service.Notify(message.Body)
                }
            }

        }
    }()

    <-forever
}