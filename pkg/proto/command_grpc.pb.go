// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: command.proto

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
	CommandProvider_AddCommand_FullMethodName        = "/proto.CommandProvider/AddCommand"
	CommandProvider_RemoveCommand_FullMethodName     = "/proto.CommandProvider/RemoveCommand"
	CommandProvider_SendCommand_FullMethodName       = "/proto.CommandProvider/SendCommand"
	CommandProvider_SendStreamCommand_FullMethodName = "/proto.CommandProvider/SendStreamCommand"
)

// CommandProviderClient is the client API for CommandProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandProviderClient interface {
	AddCommand(ctx context.Context, in *CommandInfo, opts ...grpc.CallOption) (*AddCommandResponse, error)
	RemoveCommand(ctx context.Context, in *CommandLookupRequest, opts ...grpc.CallOption) (*RemoveCommandResponse, error)
	SendCommand(ctx context.Context, in *CommandArgument, opts ...grpc.CallOption) (*CommandReturn, error)
	SendStreamCommand(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[CommandArgument, CommandReturn], error)
}

type commandProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandProviderClient(cc grpc.ClientConnInterface) CommandProviderClient {
	return &commandProviderClient{cc}
}

func (c *commandProviderClient) AddCommand(ctx context.Context, in *CommandInfo, opts ...grpc.CallOption) (*AddCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCommandResponse)
	err := c.cc.Invoke(ctx, CommandProvider_AddCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandProviderClient) RemoveCommand(ctx context.Context, in *CommandLookupRequest, opts ...grpc.CallOption) (*RemoveCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveCommandResponse)
	err := c.cc.Invoke(ctx, CommandProvider_RemoveCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandProviderClient) SendCommand(ctx context.Context, in *CommandArgument, opts ...grpc.CallOption) (*CommandReturn, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommandReturn)
	err := c.cc.Invoke(ctx, CommandProvider_SendCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandProviderClient) SendStreamCommand(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[CommandArgument, CommandReturn], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CommandProvider_ServiceDesc.Streams[0], CommandProvider_SendStreamCommand_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[CommandArgument, CommandReturn]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CommandProvider_SendStreamCommandClient = grpc.BidiStreamingClient[CommandArgument, CommandReturn]

// CommandProviderServer is the server API for CommandProvider service.
// All implementations must embed UnimplementedCommandProviderServer
// for forward compatibility.
type CommandProviderServer interface {
	AddCommand(context.Context, *CommandInfo) (*AddCommandResponse, error)
	RemoveCommand(context.Context, *CommandLookupRequest) (*RemoveCommandResponse, error)
	SendCommand(context.Context, *CommandArgument) (*CommandReturn, error)
	SendStreamCommand(grpc.BidiStreamingServer[CommandArgument, CommandReturn]) error
	mustEmbedUnimplementedCommandProviderServer()
}

// UnimplementedCommandProviderServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCommandProviderServer struct{}

func (UnimplementedCommandProviderServer) AddCommand(context.Context, *CommandInfo) (*AddCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCommand not implemented")
}
func (UnimplementedCommandProviderServer) RemoveCommand(context.Context, *CommandLookupRequest) (*RemoveCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCommand not implemented")
}
func (UnimplementedCommandProviderServer) SendCommand(context.Context, *CommandArgument) (*CommandReturn, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendCommand not implemented")
}
func (UnimplementedCommandProviderServer) SendStreamCommand(grpc.BidiStreamingServer[CommandArgument, CommandReturn]) error {
	return status.Errorf(codes.Unimplemented, "method SendStreamCommand not implemented")
}
func (UnimplementedCommandProviderServer) mustEmbedUnimplementedCommandProviderServer() {}
func (UnimplementedCommandProviderServer) testEmbeddedByValue()                         {}

// UnsafeCommandProviderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandProviderServer will
// result in compilation errors.
type UnsafeCommandProviderServer interface {
	mustEmbedUnimplementedCommandProviderServer()
}

func RegisterCommandProviderServer(s grpc.ServiceRegistrar, srv CommandProviderServer) {
	// If the following call pancis, it indicates UnimplementedCommandProviderServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CommandProvider_ServiceDesc, srv)
}

func _CommandProvider_AddCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandProviderServer).AddCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandProvider_AddCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandProviderServer).AddCommand(ctx, req.(*CommandInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandProvider_RemoveCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandLookupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandProviderServer).RemoveCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandProvider_RemoveCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandProviderServer).RemoveCommand(ctx, req.(*CommandLookupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandProvider_SendCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommandArgument)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandProviderServer).SendCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommandProvider_SendCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandProviderServer).SendCommand(ctx, req.(*CommandArgument))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandProvider_SendStreamCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandProviderServer).SendStreamCommand(&grpc.GenericServerStream[CommandArgument, CommandReturn]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CommandProvider_SendStreamCommandServer = grpc.BidiStreamingServer[CommandArgument, CommandReturn]

// CommandProvider_ServiceDesc is the grpc.ServiceDesc for CommandProvider service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandProvider_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CommandProvider",
	HandlerType: (*CommandProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCommand",
			Handler:    _CommandProvider_AddCommand_Handler,
		},
		{
			MethodName: "RemoveCommand",
			Handler:    _CommandProvider_RemoveCommand_Handler,
		},
		{
			MethodName: "SendCommand",
			Handler:    _CommandProvider_SendCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendStreamCommand",
			Handler:       _CommandProvider_SendStreamCommand_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "command.proto",
}
