package grpc

import (
	"context"
	"io"
	"strconv"

	"github.com/jasonsoft/log"

	proto "github.com/jasonsoft/grpc-example/helloworld/proto"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

// SayHello implements helloworld.GreeterServer
func (s *Server) SayHello(ctx context.Context, in *proto.HelloRequest) (*proto.HelloReply, error) {
	// md, ok := metadata.FromIncomingContext(ctx)
	// if !ok {
	// 	return nil, grpc.Errorf(codes.Unauthenticated, "無Token認證信息")
	// }

	// var (
	// 	userID string
	// 	roles  string
	// )

	// if val, ok := md["user_id"]; ok {
	// 	userID = val[0]
	// }

	// if val, ok := md["roles"]; ok {
	// 	roles = val[0]
	// }

	// log.Debugf("userID: %s roles: %s", userID, roles)
	// if userID != "jason" || roles != "admin" {
	// 	return nil, grpc.Errorf(codes.Unauthenticated, "wrong password")
	// }

	return &proto.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *Server) Check(ctx context.Context, in *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
	return &proto.HealthCheckResponse{
		Status: proto.HealthCheckResponse_SERVING,
	}, nil
}

func (s *Server) BidStream(stream proto.Chat_BidStreamServer) error {
	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			log.Info("收到客戶端通過context發出的終止信號")
			return ctx.Err()
		default:
			// 接收從客戶端發來的消息
			message, err := stream.Recv()
			if err == io.EOF {
				log.Info("客戶端發送的數據流結束")
				return nil
			}
			if err != nil {
				log.Info("接收數據出錯:", err)
				return err
			}
			// 如果接收正常，則根據接收到的 字符串 執行相應的指令
			switch message.Input {
			case "結束對話\n":
				log.Info("收到'結束對話'指令")
				if err := stream.Send(&proto.BidStreamReply{Output: "收到結束指令"}); err != nil {
					return err
				}
				// 收到結束指令時，通過 return nil 終止雙向數據流
				return nil
			case "返回數據流\n":
				log.Info("收到'返回數據流'指令")
				// 收到 收到'返回數據流'指令， 連續返回 10 條數據
				for i := 0; i < 10; i++ {
					if err := stream.Send(&proto.BidStreamReply{Output: "數據流 #" + strconv.Itoa(i)}); err != nil {
						return err
					}
				}
			default:
				// 缺省情況下， 返回 '服務端返回: ' + 輸入信息
				log.Info("[收到消息]: %s", message.Input)
				if err := stream.Send(&proto.BidStreamReply{Output: "服務端返回: " + message.Input}); err != nil {
					return err
				}
			}
		}
	}
}
