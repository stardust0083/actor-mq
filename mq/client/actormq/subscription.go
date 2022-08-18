package actormq

import "actor-mq/utils"

type Subscription struct {
	Conn        *Connection
	ChannelName string
	Ch          chan []byte
}

func (sub *Subscription) Unsubscribe() {
	context.Send(sub.Conn.ActorPid, utils.NewUnsubscribeMsg(sub.ChannelName, sub.Conn.ActorPid))
}
