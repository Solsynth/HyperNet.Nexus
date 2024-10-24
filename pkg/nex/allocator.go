package nex

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
)

type AllocatableResourceType = string

const (
	AllocatableResourceMq = AllocatableResourceType("mq")
	AllocatableResourceKv = AllocatableResourceType("kv")
)

func (v *Conn) AllocResource(t AllocatableResourceType) any {
	switch t {
	case AllocatableResourceMq:
		conn := v.GetNexusGrpcConn()
		resp, err := proto.NewAllocatorServiceClient(conn).AllocMessageQueue(context.Background(), &proto.AllocMqRequest{})
		if err != nil || !resp.IsSuccess {
			return nil
		}
		return resp.Addr
	case AllocatableResourceKv:
		conn := v.GetNexusGrpcConn()
		resp, err := proto.NewAllocatorServiceClient(conn).AllocKv(context.Background(), &proto.AllocKvRequest{})
		if err != nil || !resp.IsSuccess {
			return nil
		}
		return resp.Endpoints
	default:
		return nil
	}
}
