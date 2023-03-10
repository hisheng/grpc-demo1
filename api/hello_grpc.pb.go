// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: api/hello.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SayClient is the client API for Say service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SayClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type sayClient struct {
	cc grpc.ClientConnInterface
}

func NewSayClient(cc grpc.ClientConnInterface) SayClient {
	return &sayClient{cc}
}

func (c *sayClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/hello.Say/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayServer is the server API for Say service.
// All implementations must embed UnimplementedSayServer
// for forward compatibility
type SayServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedSayServer()
}

// UnimplementedSayServer must be embedded to have forward compatible implementations.
type UnimplementedSayServer struct {
}

func (UnimplementedSayServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedSayServer) mustEmbedUnimplementedSayServer() {}

// UnsafeSayServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SayServer will
// result in compilation errors.
type UnsafeSayServer interface {
	mustEmbedUnimplementedSayServer()
}

func RegisterSayServer(s grpc.ServiceRegistrar, srv SayServer) {
	s.RegisterService(&Say_ServiceDesc, srv)
}

func _Say_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.Say/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Say_ServiceDesc is the grpc.ServiceDesc for Say service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Say_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.Say",
	HandlerType: (*SayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Say_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/hello.proto",
}
