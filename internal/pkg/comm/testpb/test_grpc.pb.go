//
//Copyright IBM Corp. All Rights Reserved.
//
//SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: test.proto

package testpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TestService_EmptyCall_FullMethodName = "/TestService/EmptyCall"
)

// TestServiceClient is the client API for TestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TestServiceClient interface {
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type testServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTestServiceClient(cc grpc.ClientConnInterface) TestServiceClient {
	return &testServiceClient{cc}
}

func (c *testServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, TestService_EmptyCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TestServiceServer is the server API for TestService service.
// All implementations should embed UnimplementedTestServiceServer
// for forward compatibility.
type TestServiceServer interface {
	EmptyCall(context.Context, *Empty) (*Empty, error)
}

// UnimplementedTestServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTestServiceServer struct{}

func (UnimplementedTestServiceServer) EmptyCall(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCall not implemented")
}
func (UnimplementedTestServiceServer) testEmbeddedByValue() {}

// UnsafeTestServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TestServiceServer will
// result in compilation errors.
type UnsafeTestServiceServer interface {
	mustEmbedUnimplementedTestServiceServer()
}

func RegisterTestServiceServer(s grpc.ServiceRegistrar, srv TestServiceServer) {
	// If the following call pancis, it indicates UnimplementedTestServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TestService_ServiceDesc, srv)
}

func _TestService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServiceServer).EmptyCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TestService_EmptyCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServiceServer).EmptyCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// TestService_ServiceDesc is the grpc.ServiceDesc for TestService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TestService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TestService",
	HandlerType: (*TestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _TestService_EmptyCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}

const (
	EmptyService_EmptyCall_FullMethodName   = "/EmptyService/EmptyCall"
	EmptyService_EmptyStream_FullMethodName = "/EmptyService/EmptyStream"
)

// EmptyServiceClient is the client API for EmptyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmptyServiceClient interface {
	EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	EmptyStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Empty, Empty], error)
}

type emptyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmptyServiceClient(cc grpc.ClientConnInterface) EmptyServiceClient {
	return &emptyServiceClient{cc}
}

func (c *emptyServiceClient) EmptyCall(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, EmptyService_EmptyCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emptyServiceClient) EmptyStream(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[Empty, Empty], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EmptyService_ServiceDesc.Streams[0], EmptyService_EmptyStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[Empty, Empty]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EmptyService_EmptyStreamClient = grpc.BidiStreamingClient[Empty, Empty]

// EmptyServiceServer is the server API for EmptyService service.
// All implementations should embed UnimplementedEmptyServiceServer
// for forward compatibility.
type EmptyServiceServer interface {
	EmptyCall(context.Context, *Empty) (*Empty, error)
	EmptyStream(grpc.BidiStreamingServer[Empty, Empty]) error
}

// UnimplementedEmptyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmptyServiceServer struct{}

func (UnimplementedEmptyServiceServer) EmptyCall(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmptyCall not implemented")
}
func (UnimplementedEmptyServiceServer) EmptyStream(grpc.BidiStreamingServer[Empty, Empty]) error {
	return status.Errorf(codes.Unimplemented, "method EmptyStream not implemented")
}
func (UnimplementedEmptyServiceServer) testEmbeddedByValue() {}

// UnsafeEmptyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmptyServiceServer will
// result in compilation errors.
type UnsafeEmptyServiceServer interface {
	mustEmbedUnimplementedEmptyServiceServer()
}

func RegisterEmptyServiceServer(s grpc.ServiceRegistrar, srv EmptyServiceServer) {
	// If the following call pancis, it indicates UnimplementedEmptyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EmptyService_ServiceDesc, srv)
}

func _EmptyService_EmptyCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmptyServiceServer).EmptyCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmptyService_EmptyCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmptyServiceServer).EmptyCall(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmptyService_EmptyStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmptyServiceServer).EmptyStream(&grpc.GenericServerStream[Empty, Empty]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EmptyService_EmptyStreamServer = grpc.BidiStreamingServer[Empty, Empty]

// EmptyService_ServiceDesc is the grpc.ServiceDesc for EmptyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmptyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EmptyService",
	HandlerType: (*EmptyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EmptyCall",
			Handler:    _EmptyService_EmptyCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EmptyStream",
			Handler:       _EmptyService_EmptyStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}

const (
	EchoService_EchoCall_FullMethodName = "/EchoService/EchoCall"
)

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	EchoCall(ctx context.Context, in *Echo, opts ...grpc.CallOption) (*Echo, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) EchoCall(ctx context.Context, in *Echo, opts ...grpc.CallOption) (*Echo, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Echo)
	err := c.cc.Invoke(ctx, EchoService_EchoCall_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations should embed UnimplementedEchoServiceServer
// for forward compatibility.
type EchoServiceServer interface {
	EchoCall(context.Context, *Echo) (*Echo, error)
}

// UnimplementedEchoServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEchoServiceServer struct{}

func (UnimplementedEchoServiceServer) EchoCall(context.Context, *Echo) (*Echo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EchoCall not implemented")
}
func (UnimplementedEchoServiceServer) testEmbeddedByValue() {}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	// If the following call pancis, it indicates UnimplementedEchoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EchoService_ServiceDesc, srv)
}

func _EchoService_EchoCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Echo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).EchoCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EchoService_EchoCall_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).EchoCall(ctx, req.(*Echo))
	}
	return interceptor(ctx, in, info, handler)
}

// EchoService_ServiceDesc is the grpc.ServiceDesc for EchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EchoCall",
			Handler:    _EchoService_EchoCall_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}