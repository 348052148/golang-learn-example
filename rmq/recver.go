package main

import (
	"bytes"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://rbt:zh123@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	q, err := ch.QueueDeclare(
		"wannring",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	ch.QueueBind(
		q.Name,
		"xdad",
		"logs",
		false,
		nil,
	)

	msgs, err := ch.Consume(
		q.Name,
		"",
		false, //auto ack
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a constumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Printf("Done")
			d.Ack(false)
		}
	}()

	log.Printf("[*] Waiting for Message . to Exit press Ctrl + c")
	<-forever
}
