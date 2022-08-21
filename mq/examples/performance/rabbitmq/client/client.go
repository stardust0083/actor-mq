package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/streadway/amqp"
)

func openFile(fileName string) *os.File {
	f, err := os.OpenFile(fileName+".csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	return f
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	const NumberRequests = 100000

	file := openFile(fmt.Sprintf("responseTime%d", NumberRequests))
	defer file.Close()

	conn, err := amqp.Dial("amqp://guest:guest@localhost:15672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	chEncrypt, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer chEncrypt.Close()

	chEncrypted, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer chEncrypted.Close()

	msgs, err := chEncrypted.Consume(
		"encrypted", // queue
		"",          // consumer
		true,        // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)

	var initialTime time.Time

	for i := 0; i < NumberRequests; i++ {
		initialTime = time.Now()

		err = chEncrypt.Publish(
			"",        // exchange
			"encrypt", // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(fmt.Sprint(i)),
			})
		failOnError(err, "Failed to publish a message")

		msg := <-msgs

		if _, err := file.WriteString(fmt.Sprintf("%d\n", time.Since(initialTime).Nanoseconds())); err != nil {
			log.Println(err)
		}

		fmt.Println(string(msg.Body))

	}
}
