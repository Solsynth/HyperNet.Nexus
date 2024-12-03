package directory

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/kv"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"math/rand"
	"time"
)

const ServiceInfoKvPrefix = "nexus.service/"

func AddServiceInstance(in *ServiceInstance) error {
	key := ServiceInfoKvPrefix + in.ID
	data, err := json.Marshal(in)
	if err != nil {
		return err
	}

	_, err = kv.Kv.Put(context.Background(), key, string(data))
	return err
}

func GetServiceInstance(id string) *ServiceInstance {
	key := ServiceInfoKvPrefix + id
	resp, err := kv.Kv.Get(context.Background(), key)
	if err != nil || len(resp.Kvs) == 0 {
		return nil
	}

	var instance ServiceInstance
	err = json.Unmarshal(resp.Kvs[0].Value, &instance)
	if err != nil {
		return nil
	}

	return &instance
}

func ListServiceInstance() []*ServiceInstance {
	resp, err := kv.Kv.Get(context.Background(), ServiceInfoKvPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil
	}

	var result []*ServiceInstance
	for _, val := range resp.Kvs {
		var instance ServiceInstance
		if err := json.Unmarshal(val.Value, &instance); err != nil {
			continue
		}
		result = append(result, &instance)
	}
	return result
}

func ListServiceInstanceByType(t string) []*ServiceInstance {
	resp, err := kv.Kv.Get(context.Background(), ServiceInfoKvPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil
	}

	var result []*ServiceInstance
	for _, val := range resp.Kvs {
		var instance ServiceInstance
		if err := json.Unmarshal(val.Value, &instance); err != nil {
			continue
		}
		if instance.Type == t {
			result = append(result, &instance)
		}
	}
	return result
}

var srvRng = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetServiceInstanceByType(t string) *ServiceInstance {
	resp, err := kv.Kv.Get(context.Background(), ServiceInfoKvPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil
	}

	var instances []*ServiceInstance
	for _, val := range resp.Kvs {
		var instance ServiceInstance
		if err := json.Unmarshal(val.Value, &instance); err != nil {
			continue
		}
		if instance.Type == t {
			instances = append(instances, &instance)
		}
	}

	if len(instances) == 0 {
		return nil
	}

	idx := srvRng.Intn(len(instances))
	return instances[idx]
}

func RemoveServiceInstance(id string) error {
	key := ServiceInfoKvPrefix + id
	_, err := kv.Kv.Delete(context.Background(), key)
	return err
}

func BroadcastEvent(event string, data any) error {
	resp, err := kv.Kv.Get(context.Background(), ServiceInfoKvPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	log.Debug().
		Int("destinations", len(resp.Kvs)).
		Str("event", event).
		Msg("Broadcasting event from internal...")

	for idx, val := range resp.Kvs {
		var instance ServiceInstance
		if err := json.Unmarshal(val.Value, &instance); err != nil {
			log.Error().Err(err).Int("index", idx).
				Msg("Unable to parse instance config, skip broadcasting for it...")
			continue
		}

		conn, err := instance.GetGrpcConn()
		if err != nil {
			log.Error().Err(err).Str("destination", instance.ID).
				Msg("Unable to get grpc connection, skip broadcasting for it...")
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		_, _ = proto.NewDirectoryServiceClient(conn).BroadcastEvent(ctx, &proto.EventInfo{
			Event: event,
			Data:  nex.EncodeMap(data),
		})
		cancel()
	}

	return nil
}
