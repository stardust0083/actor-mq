package invocation

import (
	"actor-mq/pb"
)

type Channel struct {
	channelType pb.ChannelType
	router      *actor.PID
	subscribers map[string]*actor.PID
}

func NewChannel(channelType pb.ChannelType) *Channel {
	return &Channel{channelType: channelType, subscribers: make(map[string]*actor.PID)}
}

func (state *Channel) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *actor.Started:
		switch state.channelType {
		case pb.P2P:
			state.router = context.Spawn(router.NewRandomGroup())
		default: //plublish-subscribe
			state.router = context.Spawn(router.NewBroadcastGroup())
		}

	case *pb.SubscribeMsg:
		if subPID, ok := state.subscribers[msg.Subscriber.Address+msg.Subscriber.Id]; ok {
			// possibly reactive dormant subscriber wrapper
			context.Send(subPID, &ChangeStatus{})
		} else {
			// create new subscriber wrapper
			props := actor.PropsFromProducer(func() actor.Actor {
				return NewSubscriberWrapper(msg.Subscriber)
			})

			pid := context.Spawn(props)

			state.subscribers[msg.Subscriber.Address+msg.Subscriber.Id] = pid
			// add it to the router
			context.Send(state.router, &router.AddRoutee{pid})
		}

	case *messages.UnsubscribeMsg:
		if subPID, ok := state.subscribers[msg.Subscriber.Address+msg.Subscriber.Id]; ok {
			// remove it from router
			context.Send(state.router, &router.RemoveRoutee{subPID})
			// stop subscriber wrapper
			context.Stop(subPID)
		}
	case *messages.PublishMsg:
		if state.channelType == utils.PublishSubscribe {
			broadcastMsg := &router.BroadcastMessage{&messages.NotifyMsg{msg.ChannelName, msg.Content}}
			context.Send(state.router, broadcastMsg)
		} else { // point-to-point
			context.Send(state.router, &messages.NotifyMsg{msg.ChannelName, msg.Content})
		}
	}
}
