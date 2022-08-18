package main

import (
	"fmt"

	"actor-mq/client/actormq"
)

func main() {
	conn := actormq.Connect(actormq.ConnOptions{"127.0.0.1:8080", "127.0.0.1:8083"})

	sub := conn.Subscribe("topicx")

	for {
		msg := <-sub.Ch

		fmt.Printf("Consumer 2, received: %s", string(msg))
	}

}
