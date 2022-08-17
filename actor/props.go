package actor

type Properties interface {
	ProduceActor() Actor
	ProduceMailbox() BaseMailbox
	Supervisor() SupervisionStrategy
}

type PropsValue struct {
	actorProducer       func() Actor
	mailboxProducer     func() BaseMailbox
	supervisionStrategy SupervisionStrategy
}

func (props PropsValue) ProduceActor() Actor {
	return props.actorProducer()
}

func (props PropsValue) Supervisor() SupervisionStrategy {
	if props.supervisionStrategy == nil {
		return DefaultSupervisionStrategy()
	}
	return props.supervisionStrategy
}

func (props PropsValue) ProduceMailbox() BaseMailbox {
	if props.mailboxProducer == nil {
		return NewMailBox()
	}
	return props.mailboxProducer()
}

func Props(actorProducer func() Actor) PropsValue {
	return PropsValue{
		actorProducer:   actorProducer,
		mailboxProducer: nil,
	}
}

func (props PropsValue) WithMailbox(mailbox func() BaseMailbox) PropsValue {
	//pass by value, we only modify the copy
	props.mailboxProducer = mailbox
	return props
}

func (props PropsValue) WithSupervisor(supervisor SupervisionStrategy) PropsValue {
	//pass by value, we only modify the copy
	props.supervisionStrategy = supervisor
	return props
}
