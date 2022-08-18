package remote

import (
	"log"

	"actor-mq/actor"

	"github.com/gogo/protobuf/proto"
)

type RemoteActorRef struct {
	pid *actor.PID
}

func remoteHandler(pid *actor.PID) (actor.ActorRef, bool) {
	ref := newRemoteActorRef(pid)
	return ref, true
}

func newRemoteActorRef(pid *actor.PID) actor.ActorRef {
	return &RemoteActorRef{
		pid: pid,
	}
}

func (ref *RemoteActorRef) SendMsg(message interface{}) {
	switch msg := message.(type) {
	case proto.Message:
		envelope, _ := PackMessage(msg, ref.pid)
		endpointManagerPID.SendMsg(envelope)
	default:
		log.Printf("failed, trying to send non Proto %v message to %v", msg, ref.pid)
	}
}

func (ref *RemoteActorRef) SendCtrlMsg(message actor.SystemMessage) {}

func (ref *RemoteActorRef) Stop() {}
