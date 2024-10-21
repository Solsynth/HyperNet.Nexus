package directory

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"git.solsynth.dev/hypernet/nexus/pkg/proto"
	"sync"
	"time"
)

// In services, we use sync.Map because it will be both often read and write
var serviceDirectory sync.Map

func GetServiceInstance(id string) *ServiceInstance {
	val, ok := serviceDirectory.Load(id)
	if ok {
		return val.(*ServiceInstance)
	} else {
		return nil
	}
}

func GetServiceInstanceByType(t string) *ServiceInstance {
	var result *ServiceInstance
	serviceDirectory.Range(func(key, value any) bool {
		if value.(*ServiceInstance).Type == t {
			result = value.(*ServiceInstance)
			return false
		}
		return true
	})
	return result
}

func ListServiceInstance() []*ServiceInstance {
	var result []*ServiceInstance
	serviceDirectory.Range(func(key, value interface{}) bool {
		result = append(result, value.(*ServiceInstance))
		return true
	})
	return result
}

func ListServiceInstanceByType(t string) []*ServiceInstance {
	var result []*ServiceInstance
	serviceDirectory.Range(func(key, value interface{}) bool {
		if value.(*ServiceInstance).Type == t {
			result = append(result, value.(*ServiceInstance))
		}
		return true
	})
	return result
}

func AddServiceInstance(in *ServiceInstance) {
	serviceDirectory.Store(in.ID, in)
}

func RemoveServiceInstance(id string) {
	serviceDirectory.Delete(id)
}

func BroadcastEvent(event string, data any) {
	serviceDirectory.Range(func(key, value any) bool {
		conn, err := value.(*ServiceInstance).GetGrpcConn()
		if err != nil {
			return true
		}

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, _ = proto.NewDirectoryServiceClient(conn).BroadcastEvent(ctx, &proto.EventInfo{
			Event: event,
			Data:  nex.EncodeMap(data),
		})
		return true
	})
}
