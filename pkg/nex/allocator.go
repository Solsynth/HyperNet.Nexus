package nex

import (
	"context"

	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/redis/go-redis/v9"
)

type AllocatableResourceType = string

const (
	AllocatableResourceMq    = AllocatableResourceType("mq")
	AllocatableResourceKv    = AllocatableResourceType("kv")
	AllocatableResourceCache = AllocatableResourceType("cache")
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
	case AllocatableResourceCache:
		conn := v.GetNexusGrpcConn()
		resp, err := proto.NewAllocatorServiceClient(conn).AllocCache(context.Background(), &proto.AllocCacheRequest{})
		if err != nil || !resp.IsSuccess {
			return nil
		}
		return redis.NewClient(&redis.Options{
			Addr:     resp.GetAddr(),
			Password: resp.GetPassword(),
			DB:       int(resp.GetDb()),
		})
	default:
		return nil
	}
}
