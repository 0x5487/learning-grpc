package grpc

import (
	"context"
	"log"
	"net"
	"testing"
	"time"

	proto "github.com/jasonsoft/learning-grpc/helloworld/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"gotest.tools/assert"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

func init() {
	lis = bufconn.Listen(bufSize)
	s := grpc.NewServer()
	helloServer := NewServer()
	proto.RegisterGreeterServer(s, helloServer)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

func bufDialer(string, time.Duration) (net.Conn, error) {
	return lis.Dial()
}

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithDialer(bufDialer), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}
	defer conn.Close()
	client := proto.NewGreeterClient(conn)
	resp, err := client.SayHello(ctx, &proto.HelloRequest{Name: "Dr.Lee"})
	if err != nil {
		t.Fatalf("SayHello failed: %v", err)
	}

	assert.Equal(t, "Hello Dr.Lee", resp.Message)
}
