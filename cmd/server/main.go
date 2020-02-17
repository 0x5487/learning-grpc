package main

import (
	"net"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	helloWorldGRPC "github.com/jasonsoft/grpc-example/helloworld/delivery/grpc"
	helloworldProto "github.com/jasonsoft/grpc-example/helloworld/proto"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/log/handlers/console"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	port = ":10051"
)

// server is used to implement helloworld.GreeterServer.
func main() {
	defaultFeilds := log.Fields{
		"app_id": "grpc-server",
	}
	log.WithDefaultFields(defaultFeilds)

	clog := console.New()
	log.RegisterHandler(clog, log.AllLevels...)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    (time.Duration(5) * time.Second), // Ping the client if it is idle for 5 seconds to ensure the connection is still active
				Timeout: (time.Duration(5) * time.Second), // Wait 5 second for the ping ack before assuming the connection is dead
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             (time.Duration(2) * time.Second), // If a client pings more than once every 2 seconds, terminate the connection
				PermitWithoutStream: true,                             // Allow pings even when there are no active streams
			},
		),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			helloWorldGRPC.ErrorInterceptor(),
		)),
	)

	server := helloWorldGRPC.NewServer()
	helloworldProto.RegisterGreeterServer(s, server)
	helloworldProto.RegisterChatServer(s, server)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
