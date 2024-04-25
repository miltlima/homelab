package main

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := amqp.Dial("amqp://poc:poc@development-cluster-rabbitmq.rabbitmq.svc.cluster.local:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue %v", err)
	}

	msg := "Hello World!"

	ctx, cancel := context.WithTimeout(context.Background(), 360*time.Second)
	defer cancel()

	for {
		err = ch.PublishWithContext(ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg),
			})
		if err != nil {
			log.Fatalf("Failed to publish a message %v", err)
		} else {
			log.Printf(" [x] Sent %s", msg)
		}
		time.Sleep(300 * time.Second)
	}

}
