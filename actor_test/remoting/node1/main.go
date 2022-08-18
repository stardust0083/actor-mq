package main

import (
	"actor-mq/actor"
	"actor-mq/actor_test/remoting/messages"
	"actor-mq/remote"
	"bufio"
	"fmt"
	"log"
	"os"
)

type MyActor struct {
	count int
}

func (state *MyActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case string:
		fmt.Println(msg)
	case *messages.Response:
		state.count++
		if state.count%100 == 0 {
			log.Println(state.count)
		}
	}
}

func main() {
	remote.StartServer("localhost:8090")

	pid := actor.SpawnTemplate(&MyActor{})
	message := &messages.Echo{Message: "hej", Sender: pid}
	remoti := actor.NewPID("localhost:8091", "foo")
	for i := 0; i < 10; i++ {
		message.Message = "hej" + fmt.Sprint(i)
		remoti.SendMsg(message)
	}

	bufio.NewReader(os.Stdin).ReadString('\n')
}
