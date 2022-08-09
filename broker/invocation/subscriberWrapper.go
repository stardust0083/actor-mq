package invocation

import (
	"actor-mq/messages"

	"github.com/asynkron/protoactor-go/actor"
)

type SubscriberWrapper struct {
	remotePID      *actor.PID
	status         bool
	temporaryQueue []*messages.NotifyMsg
}

func (state *SubscriberWrapper) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		context.Watch(state.remotePID)
	case *ChangeStatus:
		state.status = true

		var x *messages.NotifyMsg
		for i := len(state.temporaryQueue) - 1; i >= 0; i-- {
			x, state.temporaryQueue = state.temporaryQueue[i], state.temporaryQueue[:i]
			context.Send(state.remotePID, x)
		}
	case *messages.NotifyMsg:
		if state.status {
			context.Send(state.remotePID, msg)
		} else {
			state.temporaryQueue = append(state.temporaryQueue, msg)
		}
	case *actor.Terminated:
		state.status = false
	}
}

func NewSubscriberWrapper(remotePID *actor.PID) *SubscriberWrapper {
	return &SubscriberWrapper{
		remotePID,
		true,
		make([]*messages.NotifyMsg, 0),
	}
}

type ChangeStatus struct{}
