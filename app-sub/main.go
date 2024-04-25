package main

import (
	"fmt"
	"log"

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

	msgs, err := ch.Consume(
		q.Name, //Name
		"",     // consumer
		true,   //AutoAck
		false,  //Exclusive
		false,  //NoLocal
		false,  //NoWait
		nil,    //Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
