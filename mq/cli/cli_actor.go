package cli

import (
	"actor-mq/actor"
	"actor-mq/mq/pb"
	"actor-mq/remote"
	"fmt"
)

type routerMgr struct {
	RouterList []*actor.PID
}

func (r *routerMgr) Receive(context actor.Context) {
	fmt.Println(context.Message())
	switch msg := context.Message().(type) {
	case *pb.SyncRouterMsg:
		r.RouterList = msg.Router
		fmt.Println(r.RouterList)
	}
}

type CliUser struct {
}

func (r *CliUser) Receive(context actor.Context) {
	fmt.Println(context.Message())
	switch msg := context.Message().(type) {
	case *actor.StateMsg:
	default:
		fmt.Println(msg)
	}
}
func StartClient(host string, port string) {
	remote.StartServer(fmt.Sprintf("%s:%s", host, port))
	pidrm := actor.SpawnTemplate(&routerMgr{RouterList: make([]*actor.PID, 0)})
	actor.PIDMgr.Register("RouterManager", pidrm)
}

func NewUser() *actor.PID {
	pid := actor.SpawnTemplate(&CliUser{})
	return pid
}

func BindUsertoRouter(localactor *actor.PID, router *actor.PID, channelName string) {
	router.SendMsg(&pb.SubscribeMsg{
		ChannelName: channelName,
		Subscriber:  localactor,
	})
}

func WriteTo(target *actor.PID, msg string) {
	target.SendMsg(&pb.CommonMsg{
		Msg:    msg,
		Target: target,
	})
}
