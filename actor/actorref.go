package actor

type ActorRef interface {
	Tell(message interface{})
	TellSystem(message SystemMessage)
	Stop()
}
type LocalActorRef struct {
	mailbox Mailbox
}

func NewLocalActorRef(mailbox Mailbox) *LocalActorRef {
	return &LocalActorRef{
		mailbox: mailbox,
	}
}

func (ref *LocalActorRef) Tell(message interface{}) {
	ref.mailbox.PostUserMessage(message)
}

func (ref *LocalActorRef) TellSystem(message SystemMessage) {
	ref.mailbox.PostSystemMessage(message)
}

func (ref *LocalActorRef) Stop() {
	ref.TellSystem(&stop{})
}

func (ref *LocalActorRef) Suspend() {
	ref.mailbox.Suspend()
}

func (ref *LocalActorRef) Resume() {
	ref.mailbox.Resume()
}

type DeadActorRef struct{}

var deadActor ActorRef = new(DeadActorRef)

func (DeadActorRef) Tell(message interface{})         {}
func (DeadActorRef) TellSystem(message SystemMessage) {}
func (DeadActorRef) Stop()                            {}
