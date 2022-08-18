package actor

type Router interface {
	Route(msg interface{})
	SetRoutee(routee []*PID)
	AddRoutee(routee []*PID)
	Routee() []*PID
}
type RouterStruct struct {
	routee []*PID
}

func (ref *RouterStruct) Route(msg interface{}) {
	for _, i := range ref.routee {
		i.SendMsg(msg)
	}
}

func (ref *RouterStruct) Routee() []*PID {
	return ref.routee
}

func (ref *RouterStruct) SetRoutee(routee []*PID) {
	ref.routee = routee
}

func (ref *RouterStruct) AddRoutee(routee []*PID) {
	ref.routee = append(ref.routee, routee...)
}
