package remote

import (
	"log"

	"actor-mq/actor"
)

var endpointManagerPID *actor.PID

func newEndpointManager() actor.Actor {
	return &endpointManager{}
}

type endpointManager struct {
	connections map[string]*actor.PID
}

func (state *endpointManager) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.StateMsg:
		switch ctx.Message().(actor.StateMsg).State {
		case actor.Started:
			state.connections = make(map[string]*actor.PID)
			log.Println("Started EndpointManager")
		}
	case *MessageEnvelope:
		pid, ok := state.connections[msg.Target.Host]
		if !ok {
			pid = actor.SpawnTemplate(&endpointWriter{host: msg.Target.Host})
			state.connections[msg.Target.Host] = pid
		}
		pid.SendMsg(msg)
	}
}
