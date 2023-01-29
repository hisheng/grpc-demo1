package main

import (
	"context"
	"fmt"
	"github/hisheng/grpc-demo1/api"
)

type hello struct {
	api.UnimplementedSayServer
}

func (h hello) SayHello(ctx context.Context, req *api.HelloRequest) (*api.HelloReply, error) {
	mes := fmt.Sprintf("来自grpc server SayHello %s", req.Name)
	return &api.HelloReply{Message: mes}, nil
}
