package actormq

import (
	"actor-mq/client/invocation"
	"actor-mq/utils"

	"actor-mq/actor"
	"actor-mq/remote"
)

var context *actor.RootContext

func init() {
	context = actor.EmptyRootContext
}

type Connection struct {
	ActorPid *actor.PID
}

func (conn *Connection) CreateChannel(channelName string, channelType utils.ChannelType) {
	context.Send(conn.ActorPid, utils.NewCreateChannelMsg(channelName, channelType))
}

func (conn *Connection) Send(channelName string, content []byte) {
	context.Send(conn.ActorPid, utils.NewPublishMsg(channelName, content))
}

func (conn *Connection) Subscribe(channelName string) *Subscription {
	ch := make(chan []byte)
	remoteMsg := utils.NewSubscribeMsg(channelName, conn.ActorPid)
	context.Send(conn.ActorPid, &invocation.SubscribeMsg{remoteMsg, ch})

	return &Subscription{conn, channelName, ch}
}

func createActor(url string) *actor.PID {

	props := actor.PropsFromProducer(func() actor.Actor {
		return &invocation.ConnectionActor{
			RemotePid:     actor.NewPID(url, "channelManager"),
			Subscriptions: make(map[string]chan []byte),
		}
	})
	return context.Spawn(props)
}

type ConnOptions struct {
	Url      string
	LocalUrl string
}

func Connect(options ConnOptions) *Connection {
	if options.LocalUrl != "" {
		remote.Start(options.LocalUrl)
	} else {
		remote.Start("")
	}

	return &Connection{createActor(options.Url)}
}
