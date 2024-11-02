// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.28.2
// source: services.proto

package proto

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

// DirectoryServiceClient is the client API for DirectoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DirectoryServiceClient interface {
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	ListService(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error)
	AddService(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*AddServiceResponse, error)
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error)
	BroadcastEvent(ctx context.Context, in *EventInfo, opts ...grpc.CallOption) (*EventResponse, error)
}

type directoryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDirectoryServiceClient(cc grpc.ClientConnInterface) DirectoryServiceClient {
	return &directoryServiceClient{cc}
}

func (c *directoryServiceClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.DirectoryService/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) ListService(ctx context.Context, in *ListServiceRequest, opts ...grpc.CallOption) (*ListServiceResponse, error) {
	out := new(ListServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.DirectoryService/ListService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) AddService(ctx context.Context, in *ServiceInfo, opts ...grpc.CallOption) (*AddServiceResponse, error) {
	out := new(AddServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.DirectoryService/AddService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error) {
	out := new(RemoveServiceResponse)
	err := c.cc.Invoke(ctx, "/proto.DirectoryService/RemoveService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *directoryServiceClient) BroadcastEvent(ctx context.Context, in *EventInfo, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/proto.DirectoryService/BroadcastEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DirectoryServiceServer is the server API for DirectoryService service.
// All implementations must embed UnimplementedDirectoryServiceServer
// for forward compatibility
type DirectoryServiceServer interface {
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	ListService(context.Context, *ListServiceRequest) (*ListServiceResponse, error)
	AddService(context.Context, *ServiceInfo) (*AddServiceResponse, error)
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)
	BroadcastEvent(context.Context, *EventInfo) (*EventResponse, error)
	mustEmbedUnimplementedDirectoryServiceServer()
}

// UnimplementedDirectoryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDirectoryServiceServer struct {
}

func (UnimplementedDirectoryServiceServer) GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetService not implemented")
}
func (UnimplementedDirectoryServiceServer) ListService(context.Context, *ListServiceRequest) (*ListServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListService not implemented")
}
func (UnimplementedDirectoryServiceServer) AddService(context.Context, *ServiceInfo) (*AddServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddService not implemented")
}
func (UnimplementedDirectoryServiceServer) RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveService not implemented")
}
func (UnimplementedDirectoryServiceServer) BroadcastEvent(context.Context, *EventInfo) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastEvent not implemented")
}
func (UnimplementedDirectoryServiceServer) mustEmbedUnimplementedDirectoryServiceServer() {}

// UnsafeDirectoryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DirectoryServiceServer will
// result in compilation errors.
type UnsafeDirectoryServiceServer interface {
	mustEmbedUnimplementedDirectoryServiceServer()
}

func RegisterDirectoryServiceServer(s grpc.ServiceRegistrar, srv DirectoryServiceServer) {
	s.RegisterService(&DirectoryService_ServiceDesc, srv)
}

func _DirectoryService_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DirectoryService/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_ListService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).ListService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DirectoryService/ListService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).ListService(ctx, req.(*ListServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_AddService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).AddService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DirectoryService/AddService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).AddService(ctx, req.(*ServiceInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DirectoryService/RemoveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DirectoryService_BroadcastEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DirectoryServiceServer).BroadcastEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DirectoryService/BroadcastEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DirectoryServiceServer).BroadcastEvent(ctx, req.(*EventInfo))
	}
	return interceptor(ctx, in, info, handler)
}

// DirectoryService_ServiceDesc is the grpc.ServiceDesc for DirectoryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DirectoryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DirectoryService",
	HandlerType: (*DirectoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetService",
			Handler:    _DirectoryService_GetService_Handler,
		},
		{
			MethodName: "ListService",
			Handler:    _DirectoryService_ListService_Handler,
		},
		{
			MethodName: "AddService",
			Handler:    _DirectoryService_AddService_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _DirectoryService_RemoveService_Handler,
		},
		{
			MethodName: "BroadcastEvent",
			Handler:    _DirectoryService_BroadcastEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
