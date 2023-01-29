package main

import (
	"context"
	"github/hisheng/grpc-demo1/api"
)

type hello struct {
	api.UnimplementedSayServer
}

func (h hello) SayHello(context.Context, *api.HelloRequest) (*api.HelloReply, error) {
	return &api.HelloReply{Message: "来自grpc server SayHello"}, nil
}
