package remote

import (
	"log"

	"actor-mq/actor"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func newEndpointWriter(host string) func() actor.Actor {
	return func() actor.Actor {
		return &endpointWriter{host: host}
	}
}

type endpointWriter struct {
	host   string
	conn   *grpc.ClientConn
	stream Remote_ReceiveMsgClient
}

func (state *endpointWriter) initialize() {
	log.Println("Started EndpointWriter for host", state.host)
	log.Println("Connecting to host", state.host)
	conn, err := grpc.Dial(state.host, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Failed to connect to host %v: %v", state.host, err)
	}
	log.Println("Connected to host", state.host)
	state.conn = conn
	c := NewRemoteClient(conn)
	log.Println("Getting stream from host", state.host)
	stream, err := c.ReceiveMsg(context.Background())
	if err != nil {
		log.Fatalf("Failed to get stream from host %v: %v", state.host, err)
	}
	log.Println("Got stream from host", state.host)
	state.stream = stream
}

func (state *endpointWriter) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.StateMsg:
		switch ctx.Message().(actor.StateMsg).State {
		case actor.Started:
			state.initialize()
		case actor.Stopped:
			state.conn.Close()
		case actor.Restarting:
			state.conn.Close()
		}
	case *MessageEnvelope:
		err := state.stream.Send(msg)
		if err != nil {
			log.Println("Failed to send to host", state.host)
			panic("restart")
		}
	default:
		log.Fatal("Unknown message", msg)
	}
}
