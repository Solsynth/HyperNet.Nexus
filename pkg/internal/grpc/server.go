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

type Server struct {
	proto.UnimplementedDatabaseControllerServer
	proto.UnimplementedStreamControllerServer
	health.UnimplementedHealthServer

	srv *grpc.Server
}

func NewServer() *Server {
	server := &Server{
		srv: grpc.NewServer(),
	}

	proto.RegisterServiceDirectoryServer(server.srv, &directory.ServiceRpcServer{})
	proto.RegisterCommandControllerServer(server.srv, &directory.CommandRpcServer{})
	proto.RegisterDatabaseControllerServer(server.srv, server)
	proto.RegisterStreamControllerServer(server.srv, server)
	health.RegisterHealthServer(server.srv, server)

	reflection.Register(server.srv)

	return server
}

func (v *Server) Listen() error {
	listener, err := net.Listen("tcp", viper.GetString("grpc_bind"))
	if err != nil {
		return err
	}

	return v.srv.Serve(listener)
}
