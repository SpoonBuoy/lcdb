// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package com

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

// StartClient is the client API for Start service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StartClient interface {
	Start(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type startClient struct {
	cc grpc.ClientConnInterface
}

func NewStartClient(cc grpc.ClientConnInterface) StartClient {
	return &startClient{cc}
}

func (c *startClient) Start(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/Start/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StartServer is the server API for Start service.
// All implementations must embed UnimplementedStartServer
// for forward compatibility
type StartServer interface {
	Start(context.Context, *Req) (*Res, error)
	mustEmbedUnimplementedStartServer()
}

// UnimplementedStartServer must be embedded to have forward compatible implementations.
type UnimplementedStartServer struct {
}

func (UnimplementedStartServer) Start(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedStartServer) mustEmbedUnimplementedStartServer() {}

// UnsafeStartServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StartServer will
// result in compilation errors.
type UnsafeStartServer interface {
	mustEmbedUnimplementedStartServer()
}

func RegisterStartServer(s grpc.ServiceRegistrar, srv StartServer) {
	s.RegisterService(&Start_ServiceDesc, srv)
}

func _Start_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StartServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Start/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StartServer).Start(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

// Start_ServiceDesc is the grpc.ServiceDesc for Start service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Start_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Start",
	HandlerType: (*StartServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _Start_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/com.proto",
}
