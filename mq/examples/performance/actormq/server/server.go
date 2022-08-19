package main

import (
	"crypto/md5"
	"fmt"

	"actor-mq/mq/ser"

	console "github.com/asynkron/goconsole"
)

func getMd5Hash(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func main() {
	const NumberRequests = 5000

	ser.StartServer("127.0.0.1", "8090")

	ser.NewRouter("encrypt")
	ser.NewRouter("encrypted")

	// for i := 0; i < NumberRequests; i++ {
	// 	msg := <-sub.Ch

	// 	conn.Send("encrypted", []byte(getMd5Hash(string(msg))))
	// }

	console.ReadLine()
}
