package actor

type ActorComm interface {
	SendUsr(message interface{})
	SendSystem(message interface{})
}

type ChannelBox struct {
	mailbox *Mailbox
}

func (box *ChannelBox) SendUsr(message interface{}) {
	box.mailbox.userMailbox <- message
	box.mailbox.schedule()
}

func (box *ChannelBox) SendSystem(message interface{}) {
	box.mailbox.systemMailbox <- message
	box.mailbox.schedule()
}
