// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: auth.proto

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
	AuthService_Authenticate_FullMethodName          = "/proto.AuthService/Authenticate"
	AuthService_EnsurePermGranted_FullMethodName     = "/proto.AuthService/EnsurePermGranted"
	AuthService_EnsureUserPermGranted_FullMethodName = "/proto.AuthService/EnsureUserPermGranted"
	AuthService_ListUserRelative_FullMethodName      = "/proto.AuthService/ListUserRelative"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error)
	EnsurePermGranted(ctx context.Context, in *CheckPermRequest, opts ...grpc.CallOption) (*CheckPermResponse, error)
	EnsureUserPermGranted(ctx context.Context, in *CheckUserPermRequest, opts ...grpc.CallOption) (*CheckUserPermResponse, error)
	ListUserRelative(ctx context.Context, in *ListUserRelativeRequest, opts ...grpc.CallOption) (*ListUserRelativeResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authenticate(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthReply)
	err := c.cc.Invoke(ctx, AuthService_Authenticate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EnsurePermGranted(ctx context.Context, in *CheckPermRequest, opts ...grpc.CallOption) (*CheckPermResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckPermResponse)
	err := c.cc.Invoke(ctx, AuthService_EnsurePermGranted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) EnsureUserPermGranted(ctx context.Context, in *CheckUserPermRequest, opts ...grpc.CallOption) (*CheckUserPermResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CheckUserPermResponse)
	err := c.cc.Invoke(ctx, AuthService_EnsureUserPermGranted_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) ListUserRelative(ctx context.Context, in *ListUserRelativeRequest, opts ...grpc.CallOption) (*ListUserRelativeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUserRelativeResponse)
	err := c.cc.Invoke(ctx, AuthService_ListUserRelative_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility.
type AuthServiceServer interface {
	Authenticate(context.Context, *AuthRequest) (*AuthReply, error)
	EnsurePermGranted(context.Context, *CheckPermRequest) (*CheckPermResponse, error)
	EnsureUserPermGranted(context.Context, *CheckUserPermRequest) (*CheckUserPermResponse, error)
	ListUserRelative(context.Context, *ListUserRelativeRequest) (*ListUserRelativeResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthServiceServer struct{}

func (UnimplementedAuthServiceServer) Authenticate(context.Context, *AuthRequest) (*AuthReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticate not implemented")
}
func (UnimplementedAuthServiceServer) EnsurePermGranted(context.Context, *CheckPermRequest) (*CheckPermResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnsurePermGranted not implemented")
}
func (UnimplementedAuthServiceServer) EnsureUserPermGranted(context.Context, *CheckUserPermRequest) (*CheckUserPermResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnsureUserPermGranted not implemented")
}
func (UnimplementedAuthServiceServer) ListUserRelative(context.Context, *ListUserRelativeRequest) (*ListUserRelativeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserRelative not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}
func (UnimplementedAuthServiceServer) testEmbeddedByValue()                     {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuthServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Authenticate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authenticate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Authenticate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authenticate(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EnsurePermGranted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EnsurePermGranted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_EnsurePermGranted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EnsurePermGranted(ctx, req.(*CheckPermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_EnsureUserPermGranted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUserPermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).EnsureUserPermGranted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_EnsureUserPermGranted_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).EnsureUserPermGranted(ctx, req.(*CheckUserPermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_ListUserRelative_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserRelativeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).ListUserRelative(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_ListUserRelative_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).ListUserRelative(ctx, req.(*ListUserRelativeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authenticate",
			Handler:    _AuthService_Authenticate_Handler,
		},
		{
			MethodName: "EnsurePermGranted",
			Handler:    _AuthService_EnsurePermGranted_Handler,
		},
		{
			MethodName: "EnsureUserPermGranted",
			Handler:    _AuthService_EnsureUserPermGranted_Handler,
		},
		{
			MethodName: "ListUserRelative",
			Handler:    _AuthService_ListUserRelative_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}
