package actor

type ActorInit func()Actor
type Actor interface{
	Receive(message Context)
}
type Receive func(Context)