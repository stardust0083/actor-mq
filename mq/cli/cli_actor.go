package cli

import (
	"actor-mq/actor"
	"actor-mq/remote"
	"fmt"
)

type CliUser struct {
}

func StartClient(host string, port string) {
	remote.StartServer(fmt.Sprintf("%s:%s", host, port))

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