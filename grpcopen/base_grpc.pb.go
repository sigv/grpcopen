// Copyright (c) 2023 Valters Jansons

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.15.8
// source: grpcopen/base.proto

package grpcopen

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

const (
	Base_Foobar_FullMethodName = "/grpopen.Base/Foobar"
	Base_Ping_FullMethodName   = "/grpopen.Base/Ping"
)

// BaseClient is the client API for Base service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BaseClient interface {
	Foobar(ctx context.Context, opts ...grpc.CallOption) (Base_FoobarClient, error)
	Ping(ctx context.Context, opts ...grpc.CallOption) (Base_PingClient, error)
}

type baseClient struct {
	cc grpc.ClientConnInterface
}

func NewBaseClient(cc grpc.ClientConnInterface) BaseClient {
	return &baseClient{cc}
}

func (c *baseClient) Foobar(ctx context.Context, opts ...grpc.CallOption) (Base_FoobarClient, error) {
	stream, err := c.cc.NewStream(ctx, &Base_ServiceDesc.Streams[0], Base_Foobar_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &baseFoobarClient{stream}
	return x, nil
}

type Base_FoobarClient interface {
	Send(*FoobarRequest) error
	Recv() (*FoobarResponse, error)
	grpc.ClientStream
}

type baseFoobarClient struct {
	grpc.ClientStream
}

func (x *baseFoobarClient) Send(m *FoobarRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *baseFoobarClient) Recv() (*FoobarResponse, error) {
	m := new(FoobarResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *baseClient) Ping(ctx context.Context, opts ...grpc.CallOption) (Base_PingClient, error) {
	stream, err := c.cc.NewStream(ctx, &Base_ServiceDesc.Streams[1], Base_Ping_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &basePingClient{stream}
	return x, nil
}

type Base_PingClient interface {
	Send(*PingRequest) error
	Recv() (*PingResponse, error)
	grpc.ClientStream
}

type basePingClient struct {
	grpc.ClientStream
}

func (x *basePingClient) Send(m *PingRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *basePingClient) Recv() (*PingResponse, error) {
	m := new(PingResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BaseServer is the server API for Base service.
// All implementations must embed UnimplementedBaseServer
// for forward compatibility
type BaseServer interface {
	Foobar(Base_FoobarServer) error
	Ping(Base_PingServer) error
	mustEmbedUnimplementedBaseServer()
}

// UnimplementedBaseServer must be embedded to have forward compatible implementations.
type UnimplementedBaseServer struct {
}

func (UnimplementedBaseServer) Foobar(Base_FoobarServer) error {
	return status.Errorf(codes.Unimplemented, "method Foobar not implemented")
}
func (UnimplementedBaseServer) Ping(Base_PingServer) error {
	return status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBaseServer) mustEmbedUnimplementedBaseServer() {}

// UnsafeBaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BaseServer will
// result in compilation errors.
type UnsafeBaseServer interface {
	mustEmbedUnimplementedBaseServer()
}

func RegisterBaseServer(s grpc.ServiceRegistrar, srv BaseServer) {
	s.RegisterService(&Base_ServiceDesc, srv)
}

func _Base_Foobar_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BaseServer).Foobar(&baseFoobarServer{stream})
}

type Base_FoobarServer interface {
	Send(*FoobarResponse) error
	Recv() (*FoobarRequest, error)
	grpc.ServerStream
}

type baseFoobarServer struct {
	grpc.ServerStream
}

func (x *baseFoobarServer) Send(m *FoobarResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *baseFoobarServer) Recv() (*FoobarRequest, error) {
	m := new(FoobarRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Base_Ping_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BaseServer).Ping(&basePingServer{stream})
}

type Base_PingServer interface {
	Send(*PingResponse) error
	Recv() (*PingRequest, error)
	grpc.ServerStream
}

type basePingServer struct {
	grpc.ServerStream
}

func (x *basePingServer) Send(m *PingResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *basePingServer) Recv() (*PingRequest, error) {
	m := new(PingRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Base_ServiceDesc is the grpc.ServiceDesc for Base service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Base_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpopen.Base",
	HandlerType: (*BaseServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Foobar",
			Handler:       _Base_Foobar_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Ping",
			Handler:       _Base_Ping_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpcopen/base.proto",
}
