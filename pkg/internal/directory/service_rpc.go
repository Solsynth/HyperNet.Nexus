package directory

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

type ServiceRpcServer struct {
	proto.UnimplementedServiceDirectoryServer
}

func instantiationService(in *ServiceInstance) *proto.ServiceInfo {
	if in == nil {
		return nil
	}
	return &proto.ServiceInfo{
		Id:       in.ID,
		Type:     in.Type,
		Label:    in.Label,
		GrpcAddr: in.GrpcAddr,
		HttpAddr: in.HttpAddr,
	}
}

func (v *ServiceRpcServer) GetService(ctx context.Context, request *proto.GetServiceRequest) (*proto.GetServiceResponse, error) {
	if request.Id != nil {
		out := GetServiceInstance(request.GetId())
		return &proto.GetServiceResponse{
			Data: instantiationService(out),
		}, nil
	}
	if request.Type != nil {
		out := GetServiceInstanceByType(request.GetType())
		return &proto.GetServiceResponse{
			Data: instantiationService(out),
		}, nil
	}
	return nil, fmt.Errorf("no filter condition is provided")
}

func (v *ServiceRpcServer) ListService(ctx context.Context, request *proto.ListServiceRequest) (*proto.ListServiceResponse, error) {
	var out []*ServiceInstance
	if request.Type != nil {
		out = ListServiceInstanceByType(request.GetType())
	} else {
		out = ListServiceInstance()
	}
	return &proto.ListServiceResponse{
		Data: lo.Map(out, func(item *ServiceInstance, index int) *proto.ServiceInfo {
			return instantiationService(item)
		}),
	}, nil
}

func (v *ServiceRpcServer) AddService(ctx context.Context, info *proto.ServiceInfo) (*proto.AddServiceResponse, error) {
	clientId, err := GetClientId(ctx)
	if err != nil {
		return nil, err
	}

	if info.GetId() != clientId {
		return nil, status.Errorf(codes.InvalidArgument, "client_id mismatch in metadata")
	}

	in := &ServiceInstance{
		ID:       clientId,
		Type:     info.GetType(),
		Label:    info.GetLabel(),
		GrpcAddr: info.GetGrpcAddr(),
		HttpAddr: info.HttpAddr,
	}
	AddServiceInstance(in)
	log.Info().Str("id", clientId).Str("label", info.GetLabel()).Msg("New service registered")
	return &proto.AddServiceResponse{
		IsSuccess: true,
	}, nil
}

func (v *ServiceRpcServer) RemoveService(ctx context.Context, request *proto.RemoveServiceRequest) (*proto.RemoveServiceResponse, error) {
	RemoveServiceInstance(request.GetId())
	log.Info().Str("id", request.GetId()).Msg("A service removed.")
	return &proto.RemoveServiceResponse{
		IsSuccess: true,
	}, nil
}
