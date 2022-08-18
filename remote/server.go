package remote

import (
	"log"
	"net"

	"actor-mq/actor"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) MsgSendRecv(stream Remote_MsgSendRecvServer) error {
	for {
		envelope, err := stream.Recv()
		if err != nil {
			return err
		}
		pid := envelope.Target
		message := UnpackMessage(envelope)
		pid.SendMsg(message)
	}
}

func StartServer(host string) {

	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	host = lis.Addr().String()
	log.Printf("Host is %v", host)
	actor.PIDMgr.RegisterHostResolver(remoteHandler)
	actor.PIDMgr.Host = host

	endpointManagerPID = actor.Spawn(actor.Props(newEndpointManager).WithMailbox(actor.NewMailBox))

	s := grpc.NewServer()
	RegisterRemoteServer(s, &server{})
	log.Printf("Starting GAM server on %v.", host)
	go s.Serve(lis)
}
