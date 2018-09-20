package main

import (
	"net"

	pb "github.com/jasonsoft/grpc-example/helloworld"
	helloWorldGRPC "github.com/jasonsoft/grpc-example/helloworld/delivery/grpc"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/log/handlers/console"
	"google.golang.org/grpc"
)

const (
	port = ":10051"
)

// server is used to implement helloworld.GreeterServer.
func main() {
	log.SetAppID("grpc") // unique id for the app

	clog := console.New()
	log.RegisterHandler(clog, log.AllLevels...)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := helloWorldGRPC.NewServer()
	pb.RegisterGreeterServer(s, server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
