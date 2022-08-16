package actor

import (
	"fmt"
	"time"
)

func RequestResponsePID() (*PID, *Response) {
	ref := &Response{
		channel: make(chan interface{}),
	}
	pid := ProcessRegistry.registerPID(ref)
	return pid, ref
}

type Response struct {
	channel chan interface{}
}

func (ref *Response) Tell(message interface{}) {
	ref.channel <- message
}

func (ref *Response) ResultChannel() <-chan interface{} {
	return ref.channel
}

func (ref *Response) ResultOrTimeout(timeout time.Duration) (interface{}, error) {
	select {
	case res := <-ref.channel:
		return res, nil
	case <-time.After(timeout):
		return nil, fmt.Errorf("Timeout")
	}
}

func (ref *Response) Result() interface{} {
	return <-ref.channel
}

func (ref *Response) TellSystem(message SystemMessage) {
}

func (ref *Response) Stop() {
	close(ref.channel)
}
