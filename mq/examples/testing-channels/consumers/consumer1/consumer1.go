package main

import (
	"fmt"

	"actor-mq/client/actormq"
)

func main() {
	conn := actormq.Connect(actormq.ConnOptions{Url: "127.0.0.1:8080", LocalUrl: "127.0.0.1:8081"})

	sub := conn.Subscribe("topicx")

	for {
		msg := <-sub.Ch

		fmt.Printf("consumer 1, received: %s", string(msg))
	}
}
