package main

import (
	"context"
	"net"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"

	"github.com/jasonsoft/log"
	"github.com/jasonsoft/log/handlers/console"
	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
)

func main() {
	log.SetAppID("grpc-proxy") // unique id for the app

	clog := console.New()
	log.RegisterHandler(clog, log.AllLevels...)

	lis, err := net.Listen("tcp", ":10080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	director := func(ctx context.Context, fullMethodName string) (context.Context, *grpc.ClientConn, error) {
		// Make sure we never forward internal services.
		log.Debugf("FullMethodName: %s", fullMethodName)
		if strings.HasPrefix(fullMethodName, "/com.example.internal.") {
			return ctx, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			val, exists := md[":authority"]
			if exists {
				log.Debugf("authority: %s", val[0])
			}

			// Decide on which backend to dial
			if val, exists := md[":authority"]; exists && val[0] == "localhost:10080" {
				// Make sure we use DialContext so the dialing can be cancelled/time out together with the context.
				ctx = metadata.NewOutgoingContext(ctx, md)
				var opts []grpc.DialOption
				opts = append(opts, grpc.WithInsecure())
				opts = append(opts, grpc.WithCodec(proxy.Codec()))
				conn, err := grpc.DialContext(ctx, "localhost:10051", opts...)
				return ctx, conn, err
			} else if val, exists := md[":authority"]; exists && val[0] == "api.example.com" {
				conn, err := grpc.DialContext(ctx, "api-service.prod.svc.local", grpc.WithCodec(proxy.Codec()))
				return ctx, conn, err
			}
		}
		return ctx, nil, grpc.Errorf(codes.Unimplemented, "Unknown method")
	}

	s := grpc.NewServer(
		grpc.CustomCodec(proxy.Codec()),
		grpc.UnknownServiceHandler(proxy.TransparentHandler(director)))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
