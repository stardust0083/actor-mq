package main

import (
	"actor-mq/actor"
	"bufio"
	"fmt"
	"os"
)

type Hello struct{ Who string }
type ParentActor struct{}

func (state *ParentActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		child := context.Spawn(actor.Props(NewChildActor))
		child.SendMsg(msg)
	}
}

func NewParentActor() actor.Actor {
	return &ParentActor{}
}

type ChildActor struct{}

func (state *ChildActor) Receive(context actor.Context) {
	msg, ok := context.Message().(actor.States)
	if ok {
		switch msg {
		case actor.States_Started:
			fmt.Println("Starting, initialize actor here")
		case actor.States_Stopping:
			fmt.Println("Stopping, actor is about shut down")
		case actor.States_Stopped:
			fmt.Println("Stopped, actor and it's children are stopped")
		case actor.States_Restarting:
			fmt.Println("Restarting, actor is about restart")
		}
	} else {
		msg := context.Message().(Hello)
		fmt.Printf("Hello %v\n", msg.Who)
		panic("Ouch")
	}
}

func NewChildActor() actor.Actor {
	return &ChildActor{}
}

func main() {
	decider := func(child *actor.PID, reason interface{}) actor.Directive {
		fmt.Println("handling failure for child")
		return actor.Directive_StopDirective
	}
	supervisor := actor.NewDefaultStrategy(10, 1000, decider)
	pid := actor.Spawn(actor.Props(NewParentActor).WithSupervisor(supervisor))
	pid.SendMsg(Hello{Who: "Roger"})

	bufio.NewReader(os.Stdin).ReadString('\n')
}
