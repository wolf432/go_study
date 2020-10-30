package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "web/proto"
)

const (
	address     = "127.0.0.1:50051"
	defaultName = "world"
)

func main() {
	//发起链接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc链接失败：%v", err)
	}
	//关闭链接
	defer conn.Close()

	//创建pb包的客户端
	client := pb.NewHelloClient(conn)

	//发起请求
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{
		Name: "linan",
	})
	if err != nil {
		log.Fatalf("recvie failed:%v", err)
	}
	fmt.Printf("value=%s", resp.Message)
}
