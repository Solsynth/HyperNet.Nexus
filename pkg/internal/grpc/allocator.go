package grpc

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/kv"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/mq"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (v *Server) AllocMessageQueue(ctx context.Context, request *proto.AllocMqRequest) (*proto.AllocMqResponse, error) {
	if mq.Kmq == nil {
		return &proto.AllocMqResponse{IsSuccess: false}, status.Error(codes.Unavailable, "message queue wasn't configured")
	}

	return &proto.AllocMqResponse{
		IsSuccess: true,
		Addr:      viper.GetString("mq.addr"),
	}, nil
}

func (v *Server) AllocKv(ctx context.Context, request *proto.AllocKvRequest) (*proto.AllocKvResponse, error) {
	if kv.Kv == nil {
		return &proto.AllocKvResponse{IsSuccess: false}, status.Error(codes.Unavailable, "kv wasn't configured")
	}

	return &proto.AllocKvResponse{
		IsSuccess: true,
		Endpoints: viper.GetStringSlice("kv.endpoints"),
	}, nil
}
