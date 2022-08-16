package actor

import "fmt"

type Context interface {
	Watch(*PID)
	UnWatch(*PID)
	Message() interface{}
	Become(Receive)
	BecomeStacked(Receive)
	UnBecomeStacked()
	Self() *PID
	Parent() *PID
	Spawn(Properties) *PID
	SpawnTemplate(Actor) *PID
	SpawnFunc(ActorInit) *PID
	Children() []*PID
}

type ContextValue struct {
	*ActorCell
	message interface{}
}

func (c *ContextValue) Message() interface{} {
	return c.message
}

func NewContext(cel *ActorCell, message interface{}) Context {
	res := &ContextValue{
		ActorCell: cel,
		message:   message,
	}
	return res
}

type ActorCell struct {
	parent     *PID
	self       *PID
	actor      Actor
	props      Properties
	supervisor SupervisionStrategy
	behaviour  []interface{}
	children   map[interface{}]struct{}
	watchers   map[interface{}]struct{}
	watching   map[interface{}]struct{}
	stopping   bool
}

func (cell *ActorCell) Children() []*PID {
	children := make([]*PID, 0)
	for i := range cell.children {
		children = append(children, i.(*PID))
	}
	return children
}

func (cell *ActorCell) Self() *PID {
	return cell.self
}

func (cell *ActorCell) Parent() *PID {
	return cell.parent
}

func NewActorCell(props Properties, parent *PID) *ActorCell {

	cell := ActorCell{
		parent:     parent,
		props:      props,
		supervisor: props.Supervisor(),
		behaviour:  make([]interface{}, 0),
		children:   make(map[interface{}]struct{}),
		watchers:   make(map[interface{}]struct{}),
		watching:   make(map[interface{}]struct{}),
	}
	cell.incarnateActor()
	return &cell
}

func (cell *ActorCell) incarnateActor() {
	actor := cell.props.ProduceActor()
	cell.actor = actor
	cell.Become(actor.Receive)
}

func (cell *ActorCell) invokeSystemMessage(message SystemMessage) {
	switch msg := message.(interface{}).(type) {
	default:
		fmt.Printf("Unknown system message %T", msg)
	case *stop:
		cell.handleStop(msg)
	case *otherStopped:
		cell.handleOtherStopped(msg)
	case *watch:
		cell.watchers[msg.Watcher] = struct{}{}
	case *unwatch:
		cell.watchers[msg.Watcher] = struct{}{}
	case *failure:
		cell.handleFailure(msg)
	case *restart:
		cell.handleRestart(msg)
	case *resume:
		cell.self.resume()
	}
}

func (cell *ActorCell) handleStop(msg *stop) {
	cell.stopping = true
	cell.invokeUserMessage(States_Stopping)
	for child := range cell.children {
		child.(*PID).Stop()
	}
	cell.tryRestartOrTerminate()
}

func (cell *ActorCell) handleOtherStopped(msg *otherStopped) {
	delete(cell.children, msg.Who)
	delete(cell.watching, msg.Who)
	cell.tryRestartOrTerminate()
}

func (cell *ActorCell) handleFailure(msg *failure) {
	directive := cell.supervisor.Handle(msg.Who, msg.Reason)
	switch directive {
	case ResumeDirective:
		//resume the fialing child
		msg.Who.TellSystem(&resume{})
	case RestartDirective:
		//restart the failing child
		msg.Who.TellSystem(&restart{})
	case StopDirective:
		//stop the failing child
		msg.Who.Stop()
	case EscalateDirective:
		//send failure to parent
		cell.parent.TellSystem(msg)
	}
}

func (cell *ActorCell) handleRestart(msg *restart) {
	cell.stopping = false
	cell.invokeUserMessage(States_Restarting) //TODO: change to restarting
	for child := range cell.children {
		child.(*PID).Stop()
	}
	cell.tryRestartOrTerminate()
}

func (cell *ActorCell) tryRestartOrTerminate() {
	if !(len(cell.children) == 0) {
		return
	}

	if !cell.stopping {
		cell.restart()
		return
	}

	cell.stopped()
}

func (cell *ActorCell) restart() {
	cell.incarnateActor()
	cell.invokeUserMessage(States_Started)
}

func (cell *ActorCell) stopped() {
	ProcessRegistry.unregisterPID(cell.self)
	cell.invokeUserMessage(States_Stopped)
	otherStopped := &otherStopped{Who: cell.self}
	for watcher := range cell.watchers {
		watcher.(*PID).TellSystem(otherStopped)
	}
}

func (cell *ActorCell) invokeUserMessage(message interface{}) {
	defer func() {
		if r := recover(); r != nil {
			failure := &failure{Reason: r, Who: cell.self}
			if cell.parent == nil {
				handleRootFailure(failure, defaultSupervisionStrategy)
			} else {
				cell.self.suspend()
				cell.parent.TellSystem(failure)
			}
		}
	}()
	if len(cell.behaviour) > 0 {
		behaviour := cell.behaviour[len(cell.behaviour)-1]
		behaviour.(Receive)(NewContext(cell, message))
	} else {
		fmt.Println("ERRRRRRRRRRRRROR", "EMpty stack")
	}

}

func (cell *ActorCell) Become(behaviour Receive) {
	cell.behaviour = make([]interface{}, 0)
	cell.behaviour = append(cell.behaviour, behaviour)
}

func (cell *ActorCell) BecomeStacked(behaviour Receive) {
	cell.behaviour = append(cell.behaviour, behaviour)
}

func (cell *ActorCell) UnbecomeStacked() {
	if len(cell.behaviour) <= 1 {
		panic("Can not unbecome actor base behaviour")
	}
	cell.behaviour = cell.behaviour[:len(cell.behaviour)-1]
}

func (cell *ActorCell) Watch(who *PID) {
	who.TellSystem(&watch{
		Watcher: cell.self,
	})
	cell.watching[who] = struct{}{}
}

func (cell *ActorCell) Unwatch(who *PID) {
	who.TellSystem(&unwatch{
		Watcher: cell.self,
	})
	delete(cell.watching, who)
}

// func (cell *ActorCell) ActorOf(props Properties) *PID {
// 	_, pid := spawnChild(props, cell.self)
// 	cell.children.Add(pid)
// 	cell.Watch(pid)
// 	return pid
// }

func (cell *ActorCell) Spawn(props Properties) *PID {
	pid := spawnChild(props, cell.self)
	cell.children[pid] = struct{}{}
	cell.Watch(pid)
	return pid
}

func (cell *ActorCell) SpawnTemplate(template Actor) *PID {
	producer := func() Actor {
		return template
	}
	props := Props(producer)
	pid := spawnChild(props, cell.self)
	cell.children[pid] = struct{}{}
	cell.Watch(pid)
	return pid
}

func (cell *ActorCell) SpawnFunc(producer ActorInit) *PID {
	props := Props(producer)
	pid := spawnChild(props, cell.self)
	cell.children[pid] = struct{}{}
	cell.Watch(pid)
	return pid
}

func handleRootFailure(msg *failure, supervisor SupervisionStrategy) {
	directive := supervisor.Handle(msg.Who, msg.Reason)
	switch directive {
	case ResumeDirective:
		//resume the fialing child
		msg.Who.TellSystem(&resume{})
	case RestartDirective:
		//restart the failing child
		msg.Who.TellSystem(&restart{})
	case StopDirective:
		//stop the failing child
		msg.Who.Stop()
	case EscalateDirective:
		//send failure to parent
		panic("Can not escalate root level failures")
	}
}
