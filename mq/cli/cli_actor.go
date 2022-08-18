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
	switch msg := context.Message().(type) {
	case *pb.SyncRouterMsg:
		r.RouterList = msg.Router
		fmt.Println(r.RouterList)
	}
}

type CliUser struct {
}

func StartClient(host string, port string) {
	remote.StartServer(fmt.Sprintf("%s:%s", host, port))
	pidrm := actor.SpawnTemplate(&routerMgr{RouterList: make([]*actor.PID, 0)})
	actor.PIDMgr.Register("RouterManager", pidrm)
}

func NewUser() *actor.PID {
	pid := actor.SpawnTemplate(&CliUser{})
}

func BindUsertoRouter(localactor *actor.PID, router *actor.PID) {

}

func WriteTo(target *actor.PID, msg interface{}) {

}

func BindUsertoRouter(localactor *actor.PID, router *actor.PID) {

}
