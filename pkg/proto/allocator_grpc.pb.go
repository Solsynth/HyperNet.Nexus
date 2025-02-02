// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: allocator.proto

package proto

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
	AllocatorService_AllocMessageQueue_FullMethodName = "/proto.AllocatorService/AllocMessageQueue"
	AllocatorService_AllocKv_FullMethodName           = "/proto.AllocatorService/AllocKv"
)

// AllocatorServiceClient is the client API for AllocatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AllocatorServiceClient interface {
	AllocMessageQueue(ctx context.Context, in *AllocMqRequest, opts ...grpc.CallOption) (*AllocMqResponse, error)
	AllocKv(ctx context.Context, in *AllocKvRequest, opts ...grpc.CallOption) (*AllocKvResponse, error)
}

type allocatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAllocatorServiceClient(cc grpc.ClientConnInterface) AllocatorServiceClient {
	return &allocatorServiceClient{cc}
}

func (c *allocatorServiceClient) AllocMessageQueue(ctx context.Context, in *AllocMqRequest, opts ...grpc.CallOption) (*AllocMqResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllocMqResponse)
	err := c.cc.Invoke(ctx, AllocatorService_AllocMessageQueue_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *allocatorServiceClient) AllocKv(ctx context.Context, in *AllocKvRequest, opts ...grpc.CallOption) (*AllocKvResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AllocKvResponse)
	err := c.cc.Invoke(ctx, AllocatorService_AllocKv_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AllocatorServiceServer is the server API for AllocatorService service.
// All implementations must embed UnimplementedAllocatorServiceServer
// for forward compatibility.
type AllocatorServiceServer interface {
	AllocMessageQueue(context.Context, *AllocMqRequest) (*AllocMqResponse, error)
	AllocKv(context.Context, *AllocKvRequest) (*AllocKvResponse, error)
	mustEmbedUnimplementedAllocatorServiceServer()
}

// UnimplementedAllocatorServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAllocatorServiceServer struct{}

func (UnimplementedAllocatorServiceServer) AllocMessageQueue(context.Context, *AllocMqRequest) (*AllocMqResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocMessageQueue not implemented")
}
func (UnimplementedAllocatorServiceServer) AllocKv(context.Context, *AllocKvRequest) (*AllocKvResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocKv not implemented")
}
func (UnimplementedAllocatorServiceServer) mustEmbedUnimplementedAllocatorServiceServer() {}
func (UnimplementedAllocatorServiceServer) testEmbeddedByValue()                          {}

// UnsafeAllocatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AllocatorServiceServer will
// result in compilation errors.
type UnsafeAllocatorServiceServer interface {
	mustEmbedUnimplementedAllocatorServiceServer()
}

func RegisterAllocatorServiceServer(s grpc.ServiceRegistrar, srv AllocatorServiceServer) {
	// If the following call pancis, it indicates UnimplementedAllocatorServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AllocatorService_ServiceDesc, srv)
}

func _AllocatorService_AllocMessageQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocMqRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocatorServiceServer).AllocMessageQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocatorService_AllocMessageQueue_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocatorServiceServer).AllocMessageQueue(ctx, req.(*AllocMqRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AllocatorService_AllocKv_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllocKvRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AllocatorServiceServer).AllocKv(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AllocatorService_AllocKv_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AllocatorServiceServer).AllocKv(ctx, req.(*AllocKvRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AllocatorService_ServiceDesc is the grpc.ServiceDesc for AllocatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AllocatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AllocatorService",
	HandlerType: (*AllocatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllocMessageQueue",
			Handler:    _AllocatorService_AllocMessageQueue_Handler,
		},
		{
			MethodName: "AllocKv",
			Handler:    _AllocatorService_AllocKv_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "allocator.proto",
}
