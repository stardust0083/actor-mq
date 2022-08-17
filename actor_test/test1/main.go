package main

import (
	"actor-mq/actor"
	"bufio"
	"fmt"
	"os"
)

type Hello struct{ Who string }
type HelloActor struct{}

func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

func main() {
	pid := actor.SpawnTemplate(&HelloActor{})
	pid.SendMsg(Hello{Who: "Roger"})
	bufio.NewReader(os.Stdin).ReadString('\n')
}
