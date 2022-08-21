package actor

import fmt "fmt"

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
	fmt.Println(msg, ref.routee)
	for _, i := range ref.routee {
		// fmt.Println(i)
		i.SendMsg(msg)
	}
}

func (ref *RouterStruct) Routee() []*PID {
	return ref.routee
}

func (ref *RouterStruct) SetRoutee(routee []*PID) {
	ref.routee = routee
}

func (ref *RouterStruct) AddRoutee(routee *PID) {
	fmt.Println("add", routee)
	for _, i := range ref.Routee() {
		if i.Host == routee.Host && i.Id == routee.Id {
			return
		}
	}
	ref.routee = append(ref.routee, routee)
}
