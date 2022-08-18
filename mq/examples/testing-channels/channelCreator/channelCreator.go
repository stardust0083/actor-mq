package main

import (
	"actor-mq/client/actormq"
	"actor-mq/pb"

	console "github.com/asynkron/goconsole"
)

func main() {
	conn := actormq.Connect(actormq.ConnOptions{Url: "127.0.0.1:8080"})

	conn.CreateChannel("topicx", pb.Subscribe)
	//conn.CreateChannel("topicx", utils.PointToPoint)

	for text, _ := console.ReadLine(); text != "exit"; text, _ = console.ReadLine() {
		conn.Send("topicx", []byte(text))
	}
}
