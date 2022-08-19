package actor

import fmt "fmt"

type RouterActorRef struct {
	ActorRef
	router *PID
	state  RouterStruct
}

func (ref *RouterActorRef) Receive(context Context) {
	switch context.Message().(type) {
	case *RouterAdd:
		ref.state.AddRoutee(context.Message().(*RouterAdd).Member)
		fmt.Println("RouterAdd", ref.state.Routee())
	case *RouterChange:
		ref.state.SetRoutee(context.Message().(*RouterChange).Members)
		fmt.Println("RouterChange", ref.state.Routee())
	case *StateMsg:

	default:
		ref.state.Route(context.Message())
	}
}

func (ref *RouterActorRef) SendMsg(message interface{}) {
	switch message.(type) {
	case *RouterAdd:
		// ref.state.AddRoutee([]*PID{message.(RouterAdd).Member})
		// fmt.Println("RouterAdd", ref.state.Routee())
	case *RouterChange:
		// ref.state.SetRoutee(message.(RouterChange).Members)
		// fmt.Println("RouterChange", ref.state.Routee())
	case *StateMsg:

	default:
		ref.state.Route(message)
	}
}

func (ref *RouterActorRef) SendCtrlMsg(message SystemMessage) {
	r, _ := PIDMgr.fromPID(ref.router)
	r.SendCtrlMsg(message)
}

func (ref *RouterActorRef) Stop() {
	ref.SendCtrlMsg(&stop{})
}

// func spawnRouter(props Props, parent *PID) *PID {
// 	id := ProcessRegistry.getAutoId()
// 	routeeProps := props
// 	routeeProps.routerConfig = nil
// 	routerState := config.Create()

// 	routerProps := FromFunc(func(context Context) {
// 		switch context.Message().(type) {
// 		case Started:
// 			config.OnStarted(context, routeeProps, routerState)
// 		}
// 	})
// 	router := Spawn(id, routerProps, parent)

// 	ref := &RouterActorRef{
// 		router: router,
// 		state:  routerState,
// 	}
// 	proxyID := PIDMgr.getAutoId()
// 	proxy := ProcessRegistry.registerPID(ref, proxyID)
// 	return proxy
// }
func NewRouterActor() Actor {
	return &RouterActorRef{}
}

func SpawnRouter() *PID {
	return SpawnTemplate(&RouterActorRef{})
}
