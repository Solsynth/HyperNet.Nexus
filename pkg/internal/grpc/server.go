package grpc

import (
	"net"

	"git.solsynth.dev/hypernet/nexus/pkg/directory"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"

	"google.golang.org/grpc/reflection"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

type GrpcServer struct {
	proto.UnimplementedStreamControllerServer

	srv *grpc.Server
}

func NewServer() *GrpcServer {
	server := &GrpcServer{
		srv: grpc.NewServer(),
	}

	proto.RegisterServiceDirectoryServer(server.srv, &directory.DirectoryRpcServer{})
	proto.RegisterStreamControllerServer(server.srv, server)
	health.RegisterHealthServer(server.srv, server)

	reflection.Register(server.srv)

	return server
}

func (v *GrpcServer) Listen() error {
	listener, err := net.Listen("tcp", viper.GetString("grpc_bind"))
	if err != nil {
		return err
	}

	return v.srv.Serve(listener)
}
