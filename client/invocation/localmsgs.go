package invocation

import "actor-mq/pb"

type SubscribeMsg struct {
	RemoteMsg *pb.SubscribeMsg
	Ch        chan []byte
}
