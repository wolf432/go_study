package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "web/proto"
)

const (
	port = ":50051"
)

// server 这个对象来实现 helloworld 包中的pb定义的rpc服务
type server struct{}

//实现接口定义中方法的具体实现
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (res *pb.HelloReply, err error) {
	return &pb.HelloReply{
		Message: "Hello" + in.Name,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}
	//生成一个rpc服务器
	s := grpc.NewServer()
	// 使用pb包调用注册已实现的rpc接口类server
	pb.RegisterHelloServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
