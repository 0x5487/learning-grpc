package main

import (
	"context"
	"log"
	"os"

	pb "github.com/jasonsoft/grpc-example/helloworld"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:10051"
	defaultName = "Jason"
)

// customCredential 自定義認證
type customCredential struct{}

func (c customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"user_id": "jason",
		"roles":   "admin",
	}, nil
}

func (c customCredential) RequireTransportSecurity() bool {
	return false
}

func main() {
	// Set up a connection to the server.
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	// 使用自定義認證
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
