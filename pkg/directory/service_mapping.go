package directory

import (
	"sync"
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
