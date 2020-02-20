package main

import (
	"context"
	"errors"
	"io"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	helloworldProto "github.com/jasonsoft/learning-grpc/helloworld/proto"
	"github.com/jasonsoft/log"
	"github.com/jasonsoft/log/handlers/console"
)

const (
	address     = "localhost:10051"
	defaultName = "jason"
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
	clog := console.New()
	log.RegisterHandler(clog, log.AllLevels...)

	conn, err := grpc.Dial(address,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(new(customCredential)), // 使用自定義認證
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                5,    // send pings every 5 seconds if there is no activity
			Timeout:             5,    // wait 5 second for ping ack before considering the connection dead
			PermitWithoutStream: true, // send pings even without active streams
		}),
		grpc.WithUnaryInterceptor(requestIDClientInterceptor()),
	)

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
		// test convert to status
		err1 := errors.New("not here")
		aaa := status.Convert(err1)
		if aaa == nil {
			log.Debug("aaa can't convert to grpc status")
		} else {
			log.Debugf("aaa to grpc status: %v", aaa.Err())
		}

		grpcStatus := status.Convert(err)
		log.Fatalf("main: could not greet: code=> %d, message => %s, ", grpcStatus.Code(), grpcStatus.Message())
	}
	log.Infof("Greeting: %s", r.Message)
}

func testChat(conn *grpc.ClientConn) {
	ctx := context.Background()

	// chat client
	client := helloworldProto.NewChatClient(conn)
	stream, err := client.BidStream(ctx)
	if err != nil {
		log.Infof("create chat client fail: [%v]\n", err)
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

func requestIDClientInterceptor() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string, req, resp interface{},
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
	) (err error) {
		logger := log.FromContext(ctx)
		logger.Debug("== begin invoker ==")

		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.Pairs()
		}

		md["request_id"] = []string{"abc-xyz"}

		err = invoker(metadata.NewOutgoingContext(ctx, md), method, req, resp, cc, opts...)
		logger.Debug("== end invoker ==")

		return
	}
}
