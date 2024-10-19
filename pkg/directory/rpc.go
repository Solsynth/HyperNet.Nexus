package directory

import (
	"context"
	"fmt"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

type DirectoryRpcServer struct {
	proto.UnimplementedServiceDirectoryServer
}

func convertServiceToInfo(in *ServiceInstance) *proto.ServiceInfo {
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

func (v *DirectoryRpcServer) GetService(ctx context.Context, request *proto.GetServiceRequest) (*proto.GetServiceResponse, error) {
	if request.Id != nil {
		out := GetServiceInstance(request.GetId())
		return &proto.GetServiceResponse{
			Data: convertServiceToInfo(out),
		}, nil
	}
	if request.Type != nil {
		out := GetServiceInstanceByType(request.GetType())
		return &proto.GetServiceResponse{
			Data: convertServiceToInfo(out),
		}, nil
	}
	return nil, fmt.Errorf("no filter condition is provided")
}

func (v *DirectoryRpcServer) ListService(ctx context.Context, request *proto.ListServiceRequest) (*proto.ListServiceResponse, error) {
	var out []*ServiceInstance
	if request.Type != nil {
		out = ListServiceInstanceByType(request.GetType())
	} else {
		out = ListServiceInstance()
	}
	return &proto.ListServiceResponse{
		Data: lo.Map(out, func(item *ServiceInstance, index int) *proto.ServiceInfo {
			return convertServiceToInfo(item)
		}),
	}, nil
}

func (v *DirectoryRpcServer) AddService(ctx context.Context, info *proto.ServiceInfo) (*proto.AddServiceResponse, error) {
	in := &ServiceInstance{
		ID:       info.GetId(),
		Type:     info.GetType(),
		Label:    info.GetLabel(),
		GrpcAddr: info.GetGrpcAddr(),
		HttpAddr: info.HttpAddr,
	}
	AddServiceInstance(in)
	log.Info().Str("id", info.GetId()).Str("label", info.GetLabel()).Msg("New service added.")
	return &proto.AddServiceResponse{
		IsSuccess: true,
	}, nil
}

func (v *DirectoryRpcServer) RemoveService(ctx context.Context, request *proto.RemoveServiceRequest) (*proto.RemoveServiceResponse, error) {
	RemoveServiceInstance(request.GetId())
	log.Info().Str("id", request.GetId()).Msg("A service removed.")
	return &proto.RemoveServiceResponse{
		IsSuccess: true,
	}, nil
}
