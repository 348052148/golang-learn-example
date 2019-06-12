package main

import (
	"log"
	"os"
	"strings"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}

//sudo docker run -d --hostname my-rabbit --name rmq -p 15672:15672 -p 5672:5672 -p 25672:25672 -e RABBITMQ_DEFAULT_USER=rbt -e RABBITMQ_DEFAULT_PASS=zh123 rabbitmq:3-management
func main() {
	//JIAN LI CONN
	conn, err := amqp.Dial("amqp://rbt:zh123@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()
	//CREATE CHANNEL
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	ch.ExchangeDeclare(
		"logs",
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)

	//CREATE QUEUE
	// q, err := ch.QueueDeclare(
	// 	"hello",
	// 	true,
	// 	false,
	// 	false,
	// 	false,
	// 	nil,
	// )
	failOnError(err, "Failed to declare a queue")
	//PUBLISH BODY
	body := "bodyFrom(os.Args1212).20"
	err = ch.Publish(
		"logs",
		"xdad",
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent, //PERSISTEN
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func bodyFrom(args []string) string {
	var s string
	if len(args) < 2 || os.Args[1] == "" {
		s = "hello"
	} else {
		strings.Join(args[:], " ")
	}
	return s
}
