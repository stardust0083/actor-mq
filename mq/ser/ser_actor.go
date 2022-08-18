package ser

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
	switch msg := context.Message().(type) {
	case *pb.CreateRouterMsg:
		if _, ok := actor.PIDMgr.LocalPids[msg.ChannelName]; ok {
			fmt.Println("Router Already Exists")
			msg.Sender.SendMsg(&pb.CreateRouterRespMsg{
				ChannelName: msg.ChannelName,
				ACK:         false,
			})
		} else {
			fmt.Println("Router Created " + msg.ChannelName)
			tmp := NewRouter(msg.ChannelName)
			r.RouterList = append(r.RouterList, tmp)
			msg.Sender.SendMsg(&pb.CreateRouterRespMsg{
				ChannelName: msg.ChannelName,
				ACK:         true,
			})
			hostManager.SendMsg(
				&syncCommand{
					RouterList: r.RouterList,
				})

		}

	}
}

var hostManager = actor.SpawnTemplate(&hostMgr{HostList: make(map[string]struct{})})

type hostMgr struct {
	HostList map[string]struct{}
}
type newHostMsg struct {
	host string
}
type syncCommand struct {
	RouterList []*actor.PID
}

func (r *hostMgr) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case newHostMsg:
		if _, ok := r.HostList[msg.host]; !ok {
			fmt.Println("New Connection " + msg.host)
			r.HostList[msg.host] = struct{}{}
		}
	case syncCommand:
		for k := range r.HostList {
			tmppid := actor.NewPID(k, "RouterManager")
			tmppid.SendMsg(
				&pb.SyncRouterMsg{
					Router: msg.RouterList,
				})
		}
	}

}

func StartServer(host string, port string) {
	remote.StartServer(fmt.Sprintf("%s:%s", host, port))
	actor.PIDMgr.Register("ConnManager", hostManager)
	pidrm := actor.SpawnTemplate(&routerMgr{RouterList: make([]*actor.PID, 0)})
	actor.PIDMgr.Register("RouterManager", pidrm)
}

func NewRouter(name string) *actor.PID {
	pidrouter := actor.SpawnTemplate(&actor.RouterActorRef{ActorRef: actor.NewActorEntity(actor.NewMailBox())})
	actor.PIDMgr.Register(name, pidrouter)
	return pidrouter
}