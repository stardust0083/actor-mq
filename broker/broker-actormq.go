package main

import (
	"actor-mq/broker/invocation"

	console "github.com/asynkron/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	rootContext := actor.EmptyRootContext

	props := actor.PropsFromProducer(func() actor.Actor { return invocation.NewChannelManager() })
	rootContext.SpawnNamed(props, "channelManager")

	console.ReadLine()
}
