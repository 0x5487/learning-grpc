package grpc

import (
	"context"

	"github.com/jasonsoft/log"

	pb "github.com/jasonsoft/grpc-example/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "無Token認證信息")
	}

	var (
		userID string
		roles  string
	)

	if val, ok := md["user_id"]; ok {
		userID = val[0]
	}

	if val, ok := md["roles"]; ok {
		roles = val[0]
	}

	log.Debugf("userID: %s roles: %s", userID, roles)
	if userID != "jason" || roles != "admin" {
		return nil, grpc.Errorf(codes.Unauthenticated, "wrong password")
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
