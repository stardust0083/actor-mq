package actor

func (pid *PID) SendMsg(message interface{}) {
	ref, _ := PIDMgr.fromPID(pid)
	ref.SendMsg(message)
}

func SendMsg(pid *PID, message interface{}) {
	ref, _ := PIDMgr.fromPID(pid)
	ref.SendMsg(message)
}

func (pid *PID) SendCtrlMsg(message SystemMessage) {
	ref, _ := PIDMgr.fromPID(pid)
	ref.SendCtrlMsg(message)
}

func (pid *PID) Stop() {
	ref, _ := PIDMgr.fromPID(pid)
	ref.Stop()
}

func (pid *PID) suspend() {
	ref, _ := PIDMgr.fromPID(pid)
	ref.(*ActorEntity).Suspend()
}

func (pid *PID) resume() {
	ref, _ := PIDMgr.fromPID(pid)
	ref.(*ActorEntity).Resume()
}

func NewPID(host, id string) *PID {
	return &PID{
		Host: host,
		Id:   id,
	}
}
