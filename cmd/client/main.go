package main

import (
	"google.golang.org/grpc/status"
	"context"
	"io"
	"os"
	"time"

	helloworldProto "github.com/jasonsoft/grpc-example/helloworld/proto"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/log/handlers/console"
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
	ctx := context.Background()

	log.SetAppID("grpc-server") // unique id for the app

	clog := console.New()
	log.RegisterHandler(clog, log.AllLevels...)

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

	c := helloworldProto.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &helloworldProto.HelloRequest{Name: name})
	if err != nil {
		s := status.Convert(err)
		log.Fatalf("main: could not greet: code=> %d, message => %s, ", s.Code() , s.Message())
	}
	log.Infof("Greeting: %s", r.Message)

	// health check
	healthCheckRequest := &helloworldProto.HealthCheckRequest{
		Service: "aaa",
	}
	healthCheckResp, err := c.Check(ctx, healthCheckRequest)
	if err != nil {
		log.Errorf("Error: %v", err)
	}

	log.Debug("healthCheckResp : %s", healthCheckResp.Status)
}

func testPing(conn *grpc.ClientConn) {
	ctx := context.Background()

	client := helloworldProto.NewPingPongClient(conn)

	pingRequest := &helloworldProto.PingRequest{Message: "Ping"}
	pongReply, err := client.Ping(ctx, pingRequest)
	if err != nil {
		log.Errorf("test ping err: %v", err)
	}
	log.Debugf(pongReply.Message)
}

func testChat(conn *grpc.ClientConn) {
	ctx := context.Background()

	// chat client
	client := helloworldProto.NewChatClient(conn)
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Info("create chat client fail: [%v]\n", err)
	}

	// send message to server
	go func() {
		for {
			err := stream.Send(&helloworldProto.BidStreamRequest{Input: "你好"})
			if err != nil {
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			log.Info("收到服務端的結束信號")
			break //如果收到結束信號，則退出“接收循環”，結束客戶端程序
		}
		if err != nil {
			// TODO: 處理接收錯誤
			log.Errorf("接收數據出錯: %v", err)
		}
		// 沒有錯誤的情況下，打印來自服務端的消息
		log.Infof("[客戶端收到]: %s", message.Output)
	}
}
