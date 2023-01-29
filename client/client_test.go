package client

import (
	"context"
	"github/hisheng/grpc-demo1/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"testing"
)

func TestSayHello(t *testing.T) {
	// 1、初始化grpc连接conn
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	serverAddr := "localhost:8083"
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 2、初始化sayClient
	client := api.NewSayClient(conn)
	// 3、调用sayClient里面的方法
	resp, err := client.SayHello(context.Background(), &api.HelloRequest{Name: "海生"})
	log.Fatal(resp, err)
}
