package main

import (
	"github/hisheng/grpc-demo1/api"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	// 1.8083 作为grpc的监听端口号
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 2. 新建一个 grpc server
	server := grpc.NewServer()
	// 3 . 注册自己的方法 到 grpc服务器中
	// todo 使用 _grpc.pb.go 里面的register方法
	api.RegisterSayServer(server, hello{})

	// 4. 启动 grpc server
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// 5.关闭 grpc server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
