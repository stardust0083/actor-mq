package actor

type Actor interface {
	Receive(message Context)
}
type ActorRef interface {
	SendMsg(message interface{})
	SendCtrlMsg(message SystemMessage)
	Stop()
}
type ActorEntity struct {
	mailbox BaseMailbox
}

func NewActorEntity(mailbox BaseMailbox) *ActorEntity {
	return &ActorEntity{mailbox: mailbox}
}

func (ref *ActorEntity) SendMsg(message interface{}) {
	ref.mailbox.PostUserMessage(message)
}

func (ref *ActorEntity) SendCtrlMsg(message SystemMessage) {
	ref.mailbox.PostSystemMessage(message)
}

func (ref *ActorEntity) Stop() {
	ref.SendCtrlMsg(&stop{})
}

func (ref *ActorEntity) Suspend() {
	ref.mailbox.Suspend()
}

func (ref *ActorEntity) Resume() {
	ref.mailbox.Resume()
}

type EmptyActorEntity struct{}

var emptyActor ActorRef = new(EmptyActorEntity)

func (EmptyActorEntity) SendMsg(message interface{})       {}
func (EmptyActorEntity) SendCtrlMsg(message SystemMessage) {}
func (EmptyActorEntity) Stop()                             {}
