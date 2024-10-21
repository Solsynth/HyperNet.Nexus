package grpc

import (
	directory2 "git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"net"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"

	"google.golang.org/grpc/reflection"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	proto.UnimplementedDatabaseServiceServer
	proto.UnimplementedStreamServiceServer
	health.UnimplementedHealthServer

	srv *grpc.Server
}

func NewServer() *Server {
	server := &Server{
		srv: grpc.NewServer(),
	}

	proto.RegisterDirectoryServiceServer(server.srv, &directory2.ServiceRpcServer{})
	proto.RegisterCommandProviderServer(server.srv, &directory2.CommandRpcServer{})
	proto.RegisterDatabaseServiceServer(server.srv, server)
	proto.RegisterStreamServiceServer(server.srv, server)
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
