package main

import (
	"crypto/md5"
	"fmt"

	"actor-mq/client/actormq"
	"actor-mq/pb"

	console "github.com/asynkron/goconsole"
)

func getMd5Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func main() {
	const NumberRequests = 5000

	conn := actormq.Connect(actormq.ConnOptions{Url: "127.0.0.1:8080", LocalUrl: "127.0.0.1:8081"})

	conn.CreateChannel("encrypt", pb.Subscribe)
	conn.CreateChannel("encrypted", pb.Subscribe)

	sub := conn.Subscribe("encrypt")
	//defer sub.Unsubscribe()

	for i := 0; i < NumberRequests; i++ {
		msg := <-sub.Ch

		conn.Send("encrypted", []byte(getMd5Hash(string(msg))))
	}

	console.ReadLine()
	sub.Unsubscribe()
}
