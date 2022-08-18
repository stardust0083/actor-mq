package main

import (
	"fmt"
	"time"

	"actor-mq/actor"

	console "github.com/asynkron/goconsole"
)

type EchoActor struct{}

func NewEchoActor() actor.Actor {
	return &EchoActor{}
}

type BlackHoleActor struct{}

func (state *BlackHoleActor) Receive(context actor.Context) {
	fmt.Println("Cavalry has arrived", context.Message())
}

func NewBlackHoleActor() actor.Actor {
	return &BlackHoleActor{}
}

func (*EchoActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case EchoMessage:
		fmt.Println("route")
		msg.Router.SendMsg(EchoReplyMessage{context.Message().(EchoMessage).num})
	}
}

type EchoMessage struct {
	num    int
	Router *actor.PID
}

type EchoReplyMessage struct{ num int }

func main() {
	pidorigin := actor.SpawnTemplate(&EchoActor{})
	actor.PIDMgr.Register("origin", pidorigin)
	pidrouter := actor.SpawnTemplate(&actor.RouterActorRef{ActorRef: actor.NewActorEntity(actor.NewMailBox())})
	actor.PIDMgr.Register("router", pidrouter)
	pidblack1 := actor.SpawnTemplate(&BlackHoleActor{})
	actor.PIDMgr.Register("black1", pidblack1)
	pidblack2 := actor.SpawnTemplate(&BlackHoleActor{})
	actor.PIDMgr.Register("black2", pidblack2)
	pidblack3 := actor.SpawnTemplate(&BlackHoleActor{})
	actor.PIDMgr.Register("black3", pidblack3)
	pidrouter.SendMsg(actor.RouterAdd{Member: pidblack1})

	for i := 0; i < 5; i++ {
		pidorigin.SendMsg(EchoMessage{i, pidrouter})
	}
	time.Sleep(1 * time.Second)
	pidrouter.SendMsg(actor.RouterChange{Members: []*actor.PID{pidblack2, pidblack3}})
	for i := 6; i < 10; i++ {
		pidorigin.SendMsg(EchoMessage{i, pidrouter})
	}
	time.Sleep(1 * time.Second)
	pidrouter.SendMsg(actor.RouterAdd{Member: pidblack1})
	for i := 11; i < 15; i++ {
		pidorigin.SendMsg(EchoMessage{i, pidrouter})
	}
	console.ReadLine()
}
