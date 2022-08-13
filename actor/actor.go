package actor

import "fmt"

type BaseActor interface {
	Receive(message *Context)
}

type LocalActorCore struct {
	Comm     ActorComm
	actor    BaseActor
	behavior func(*Context)
}

func NewLocalActorCore(actor BaseActor) *LocalActorCore {
	core := LocalActorCore{
		actor:    actor,
		behavior: actor.Receive,
	}
	return &core
}

type Context struct {
	owner *LocalActorCore
	Msg   interface{}
}

func InitActor(actor BaseActor) ActorComm {
	userMailbox := make(chan interface{}, 100)
	systemMailbox := make(chan interface{}, 100)
	core := NewLocalActorCore(actor)
	mailbox := Mailbox{
		userMailbox:     userMailbox,
		systemMailbox:   systemMailbox,
		hasMoreMessages: MailboxHasNoMessages,
		schedulerStatus: MailboxIdle,
		parent:          core,
	}
	cb := ChannelBox{mailbox: &mailbox}
	core.Comm = &cb
	return &cb
}

func (core *LocalActorCore) invokeSystemMessage(message interface{}) {
	fmt.Printf("Received system message %v\n", message)
}

func (core *LocalActorCore) invokeUserMessage(message interface{}) {
	context := Context{
		owner: core,
		Msg:   message,
	}
	core.behavior(&context)
}

func (core *LocalActorCore) Become(behavior func(*Context)) {
	core.behavior = behavior
}
