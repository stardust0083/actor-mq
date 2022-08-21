package main

import (
	"crypto/md5"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func getMd5Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:15672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	chEncrypt, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer chEncrypt.Close()

	chEncrypted, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer chEncrypted.Close()

	qEncrypt, err := chEncrypt.QueueDeclare(
		"encrypt", // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	failOnError(err, "Failed to declare a queue")

	qEncrypted, err := chEncrypted.QueueDeclare(
		"encrypted", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := chEncrypt.Consume(
		qEncrypt.Name, // queue
		"",            // consumer
		true,          // auto-ack
		false,         // exclusive
		false,         // no-local
		false,         // no-wait
		nil,           // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			chEncrypted.Publish(
				"",              // exchange
				qEncrypted.Name, // routing key
				false,           // mandatory
				false,           // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        d.Body,
				})
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}
