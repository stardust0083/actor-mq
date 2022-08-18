package main

import (
	"actor-mq/actor"
	"actor-mq/actor_test/remoting/messages"
	"actor-mq/remote"
	"bufio"
	"fmt"
	"os"
)

type MyActor struct{}

func (*MyActor) Receive(context actor.Context) {
	fmt.Println(context.Self())
	switch msg := context.Message().(type) {
	case *messages.Echo:
		fmt.Println("send")
		msg.Sender.SendMsg(&messages.Response{
			SomeValue: "result",
			AnInt:     123,
		})
	}
}

func main() {
	remote.StartServer("localhost:8091")
	pid := actor.SpawnTemplate(&MyActor{})
	actor.PIDMgr.Register("foo", pid)
	bufio.NewReader(os.Stdin).ReadString('\n')
}
