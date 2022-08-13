package actor

import "sync/atomic"

const (
	MailboxIdle    int32 = 0
	MailboxRunning int32 = 1
)
const (
	MailboxHasNoMessages   int32 = 0
	MailboxHasMoreMessages int32 = 1
)

type Mailbox struct {
	parent          *LocalActorCore
	userMailbox     chan interface{}
	systemMailbox   chan interface{}
	schedulerStatus int32
	hasMoreMessages int32
}

func (mailbox *Mailbox) schedule() {
	swapped := atomic.CompareAndSwapInt32(&mailbox.schedulerStatus, MailboxIdle, MailboxRunning)
	atomic.StoreInt32(&mailbox.hasMoreMessages, MailboxHasMoreMessages) //we have more messages to process
	if swapped {
		go mailbox.processMessages()
	}
}

func (mailbox *Mailbox) processMessages() {
	atomic.StoreInt32(&mailbox.hasMoreMessages, MailboxHasNoMessages)
	for i := 0; i < 30; i++ {
		select {
		case sysMsg := <-mailbox.systemMailbox:
			mailbox.parent.invokeSystemMessage(sysMsg)
		default:
			select {
			case userMsg := <-mailbox.userMailbox:
				mailbox.parent.invokeUserMessage(userMsg)
			default:
			}
		}
	}
	atomic.StoreInt32(&mailbox.schedulerStatus, MailboxIdle)
	hasMore := atomic.LoadInt32(&mailbox.hasMoreMessages)
	status := atomic.LoadInt32(&mailbox.schedulerStatus)
	if hasMore == MailboxHasMoreMessages && status == MailboxIdle {
		swapped := atomic.CompareAndSwapInt32(&mailbox.schedulerStatus, MailboxIdle, MailboxRunning)
		if swapped {
			go mailbox.processMessages()
		}
	}
}
