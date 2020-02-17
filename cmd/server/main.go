package main

import (
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	helloWorldGRPC "github.com/jasonsoft/grpc-example/helloworld/delivery/grpc"
	helloworldProto "github.com/jasonsoft/grpc-example/helloworld/proto"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/log/handlers/console"
	"google.golang.org/grpc"
)

const (
	port = ":10051"
)

// server is used to implement helloworld.GreeterServer.
func main() {
	log.SetAppID("grpc-server") // unique id for the app

	clog := console.New()
	log.RegisterHandler(clog, log.AllLevels...)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			helloWorldGRPC.UnaryServerInterceptor(),
		)),
	)

	server := helloWorldGRPC.NewServer()
	helloworldProto.RegisterGreeterServer(s, server)
	helloworldProto.RegisterChatServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
