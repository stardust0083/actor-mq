package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"actor-mq/client/actormq"
)

func openFile(fileName string) *os.File {
	f, err := os.OpenFile(fileName+".csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}

	return f
}

func main() {
	const NumberRequests = 5000

	file := openFile(fmt.Sprintf("responseTime%d", NumberRequests))
	defer file.Close()

	conn := actormq.Connect(actormq.ConnOptions{Url: "127.0.0.1:8080", LocalUrl: "127.0.0.1:8082"})

	sub := conn.Subscribe("encrypted")
	defer sub.Unsubscribe()

	var initialTime time.Time

	for i := 0; i < NumberRequests; i++ {
		initialTime = time.Now()

		conn.Send("encrypt", []byte("Hello World!"))

		msg := <-sub.Ch

		if _, err := file.WriteString(fmt.Sprintf("%d\n", time.Since(initialTime).Nanoseconds())); err != nil {
			log.Println(err)
		}

		fmt.Println(string(msg))
	}
}
