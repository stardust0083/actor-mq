package invocation

import (
	"fmt"
	"log"

	"actor-mq/messages"
	"actor-mq/utils"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/remote"
)

const host = "localhost"
const port = 8080

func init() {
	fmt.Printf("[INFO] Broker starting at %s:%d\n", host, port)
	remote.Start(fmt.Sprintf("%s:%d", host, port))
}

type ChannelManager struct {
	channels map[string]*actor.PID
}

func NewChannelManager() *ChannelManager {
	return &ChannelManager{make(map[string]*actor.PID)}
}

func (state *ChannelManager) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case *messages.CreateChannelMsg:
		log.Println("[INFO] Received Create Channel Msg")

		if _, ok := state.channels[msg.ChannelName]; !ok {
			props := actor.PropsFromProducer(func() actor.Actor {
				return NewChannel(utils.ChannelType(msg.ChannelType))
			})

			childPid := context.Spawn(props)

			state.channels[msg.ChannelName] = childPid

			log.Printf("[INFO] Channel %s Created", msg.ChannelName)
		}

	case *messages.PublishMsg:
		log.Println("[INFO] Received Send Msg")

		if channelPID, ok := state.channels[msg.ChannelName]; ok {
			context.Send(channelPID, msg)
		}
	case *messages.SubscribeMsg:
		log.Println("[INFO] Received Subscribe Msg")

		if channelPID, ok := state.channels[msg.ChannelName]; ok {
			context.Send(channelPID, msg)
		}
	case *messages.UnsubscribeMsg:
		log.Println("[INFO] Received Unsubscribe Msg")

		if channelPID, ok := state.channels[msg.ChannelName]; ok {
			context.Send(channelPID, msg)
		}
	}
}
