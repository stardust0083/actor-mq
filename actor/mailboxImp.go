package actor

import (
	"sync/atomic"
)

const chanmaxlen = 100
const (
	MailboxIdle    int32 = 0
	MailboxRunning int32 = 1
)
const (
	MailboxHasNoMessages   int32 = 0
	MailboxHasMoreMessages int32 = 1
)

type MailboxProducer func() BaseMailbox
type BaseMailbox interface {
	PostUserMessage(message interface{})
	PostSystemMessage(message SystemMessage)
	Suspend()
	Resume()
	RegisterHandlers(userInvoke func(interface{}), systemInvoke func(SystemMessage))
}

type Mailbox struct {
	userMailbox     chan interface{}
	systemMailbox   chan SystemMessage
	schedulerStatus int32
	hasMoreMessages int32
	userInvoke      func(interface{})
	systemInvoke    func(SystemMessage)
}

func (mb *Mailbox) PostUserMessage(message interface{}) {
	mb.userMailbox <- message
	mb.schedule()
}

func (mb *Mailbox) PostSystemMessage(message SystemMessage) {
	mb.userMailbox <- message
	mb.schedule()
}

func (mb *Mailbox) schedule() {
	atomic.StoreInt32(&mb.hasMoreMessages, MailboxHasMoreMessages) //we have more messages to process
	if atomic.CompareAndSwapInt32(&mb.schedulerStatus, MailboxIdle, MailboxRunning) {
		go mb.processMessages()
	}
}

func (mb *Mailbox) Suspend() {

}

func (mb *Mailbox) Resume() {

}

func (mb *Mailbox) processMessages() {
	atomic.StoreInt32(&mb.hasMoreMessages, MailboxHasMoreMessages)
	done := false
	for i := 0; i < chanmaxlen; i++ {
		select {
		case sysMsg := <-mb.systemMailbox:
			mb.systemInvoke(sysMsg)
		default:
			select {
			case Msg := <-mb.userMailbox:
				mb.userInvoke(Msg)
			default:
				done = true
				break
			}
		}
	}
	if !done {
		atomic.StoreInt32(&mb.hasMoreMessages, MailboxHasMoreMessages)
	}
	atomic.StoreInt32(&mb.schedulerStatus, MailboxIdle)
	if atomic.SwapInt32(&mb.hasMoreMessages, MailboxHasNoMessages) == MailboxHasMoreMessages {
		mb.schedule()
	}
}

func (mb *Mailbox) RegisterHandlers(usr func(interface{}), sys func(SystemMessage)) {
	mb.userInvoke = usr
	mb.systemInvoke = sys
}

func NewMailBox() BaseMailbox {
	usrMailbox := make(chan interface{}, chanmaxlen)
	systemMailbox := make(chan SystemMessage, chanmaxlen)
	mailbox := Mailbox{
		userMailbox:     usrMailbox,
		systemMailbox:   systemMailbox,
		hasMoreMessages: MailboxHasNoMessages,
		schedulerStatus: MailboxIdle,
	}
	return &mailbox
}
