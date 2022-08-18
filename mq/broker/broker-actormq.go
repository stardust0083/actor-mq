package main

import (
	"actor-mq/broker/invocation"

	"actor-mq/actor"

	console "github.com/asynkron/goconsole"
)

func main() {
	rootContext := actor.EmptyRootContext

	props := actor.PropsFromProducer(func() actor.Actor { return invocation.NewChannelManager() })
	rootContext.SpawnNamed(props, "channelManager")

	console.ReadLine()
}
