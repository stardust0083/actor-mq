package actor

func SpawnFunc(producer func() Actor) *PID {
	props := Props(producer)
	pid := spawnChild(props, nil)
	return pid
}

func SpawnTemplate(template Actor) *PID {
	//actorType := reflect.TypeOf(template)
	producer := func() Actor {
		//	return reflect.New(actorType).Elem().Interface().(Actor)
		return template
	}
	props := Props(producer)
	pid := spawnChild(props, nil)
	return pid
}

func Spawn(props Properties) *PID {
	pid := spawnChild(props, nil)
	return pid
}

func spawnChild(props Properties, parent *PID) *PID {
	cell := NewActorCell(props, parent)
	mailbox := props.ProduceMailbox()
	mailbox.RegisterHandlers(cell.invokeUserMessage, cell.invokeSystemMessage)
	ref := NewActorEntity(mailbox)
	pid := PIDMgr.registerPID(ref)
	cell.self = pid
	cell.invokeUserMessage(States_Started)
	return pid
}
