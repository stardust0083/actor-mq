package cli

import (
	"actor-mq/actor"
	"actor-mq/mq/pb"
	"actor-mq/remote"
	"fmt"
	"os"
	"time"
)

var f *os.File

func init() {
	f, _ = os.OpenFile("respondtime.csv", os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
}

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
	switch msg := context.Message().(type) {
	case *actor.StateMsg:
	case *pb.CommonMsg:
		t1, _ := time.Parse(`2006-01-02T15:04:05.000000000`, msg.Time)

		fmt.Fprintln(f, time.Since(t1.Add(-time.Hour*8)).Nanoseconds())
		fmt.Println("Receive", msg.Msg, time.Since(t1.Add(-time.Hour*8)))
	default:
		// fmt.Println(msg, time.Now())
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

func BindUsertoRouter(localactor *actor.PID, router *actor.PID) {
	router.SendMsg(&actor.RouterAdd{
		Member: localactor,
	})
}

func WriteTo(target *actor.PID, msg string) {
	target.SendMsg(&pb.CommonMsg{
		Msg:    msg,
		Target: target,
		Time:   time.Now().Format(`2006-01-02T15:04:05.000000000`),
	})
}

func CloseFile() {
	f.Close()
}
