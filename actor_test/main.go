package main

import (
	"bufio"
	"fmt"
	"os"

	"actor-mq/actor"
)

func main() {
	parent := actor.InitActor(NewParentActor())
	parent.SendUsr(Hello{Name: "Roger"})
	parent.SendUsr(Hello{Name: "Go"})
	bufio.NewReader(os.Stdin).ReadString('\n')
}

type Ping struct {
	Sender actor.LocalActorCore
	Name   string
}
type Pong struct{}
type Hello struct{ Name string }

type ChildActor struct{ messageCount int }

func (state *ChildActor) Receive(context *actor.Context) {
	switch msg := context.Msg.(type) {
	default:
		fmt.Printf("unexpected type %T\n", msg)
	case Ping:
		fmt.Printf("Hello %v\n", msg.Name)
		state.messageCount++
		fmt.Printf("message count %v \n", state.messageCount)
		msg.Sender.Comm.SendUsr(Pong{})
	}
}

func NewParentActor() actor.BaseActor {
	return &ParentActor{
		Child: actor.InitActor(new(ChildActor)),
	}
}

type ParentActor struct {
	Child actor.LocalActorCore
}

func (state *ParentActor) Receive(context *actor.Context) {
	switch msg := context.Message.(type) {
	default:
		fmt.Printf("unexpected type %T\n", msg)
	case Hello:
		fmt.Printf("Parent got hello %v\n", msg.Name)
		state.Child.Tell(Ping{
			Name:   msg.Name,
			Sender: context.Self,
		})
		context.Become(state.Other)
	}
}

func (state *ParentActor) Other(context *actor.Context) {
	switch msg := context.Message.(type) {
	default:
		fmt.Printf("unexpected type %T\n", msg)
	case Pong:
		fmt.Println("Got pong")
	}
}
