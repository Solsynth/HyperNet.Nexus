package grpc

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/database"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (v *Server) AllocDatabase(ctx context.Context, request *proto.AllocDatabaseRequest) (*proto.AllocDatabaseResponse, error) {
	dsn, err := database.AllocDatabase(request.GetName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &proto.AllocDatabaseResponse{
		IsSuccess: true,
		Dsn:       dsn,
	}, nil
}
