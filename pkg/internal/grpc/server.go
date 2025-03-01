package grpc

import (
	"net"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"

	"google.golang.org/grpc/reflection"

	"github.com/spf13/viper"
	"google.golang.org/grpc"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	proto.UnimplementedDatabaseServiceServer
	proto.UnimplementedStreamServiceServer
	proto.UnimplementedAllocatorServiceServer
	health.UnimplementedHealthServer

	srv *grpc.Server
}

func NewServer() *Server {
	server := &Server{
		srv: grpc.NewServer(),
	}

	proto.RegisterDirectoryServiceServer(server.srv, &directory.ServiceRpcServer{})
	proto.RegisterDatabaseServiceServer(server.srv, server)
	proto.RegisterStreamServiceServer(server.srv, server)
	proto.RegisterAllocatorServiceServer(server.srv, server)
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
