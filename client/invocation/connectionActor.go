package invocation

import (
	"log"

	"actor-mq/pb"

	"actor-mq/actor"
)

type ConnectionActor struct {
	RemotePid     *actor.PID
	Subscriptions map[string]chan []byte
}

func (state *ConnectionActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *pb.CreateChannelMsg:
		log.Println("[INFO] Received Create Channel Message")

		context.Send(state.RemotePid, msg)

	case *pb.PublishMsg:
		log.Println("[INFO] Received Publish Msg")

		context.Send(state.RemotePid, msg)
	case *pb.NotifyMsg:
		log.Println("[INFO] Received Notify Msg")

		if ch, ok := state.Subscriptions[msg.Name]; ok {
			ch <- msg.Content
		}

	case *pb.SubscribeMsg:
		log.Println("[INFO] Received Subscribe Msg")

		state.Subscriptions[msg.RemoteMsg.Name] = msg.Ch
		context.Send(state.RemotePid, msg.RemoteMsg)

	case *pb.UnsubscribeMsg:
		log.Println("[INFO] Received Unsubscribe Msg")

		delete(state.Subscriptions, msg.ChannelName)
		context.Send(state.RemotePid, msg)
	}
}
