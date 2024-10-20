package directory

import (
	"github.com/samber/lo"
	"sync"
)

// In commands, we use the map and the mutex because it is usually read and only sometimes write
var commandDirectory = make(map[string]*Command)
var commandDirectoryMutex sync.Mutex

func GetCommandKey(id, method string) string {
	return id + ":" + method
}

func AddCommand(id, method string, tags []string, handler *ServiceInstance) {
	commandDirectoryMutex.Lock()
	defer commandDirectoryMutex.Unlock()

	ky := GetCommandKey(id, method)
	if _, ok := commandDirectory[id]; !ok {
		commandDirectory[id] = &Command{
			ID:      id,
			Method:  method,
			Tags:    tags,
			Handler: []*ServiceInstance{handler},
		}
	} else {
		commandDirectory[ky].Handler = append(commandDirectory[ky].Handler, handler)
		commandDirectory[ky].Tags = lo.Uniq(append(commandDirectory[ky].Tags, tags...))
	}

	commandDirectory[ky].Handler = lo.UniqBy(commandDirectory[ky].Handler, func(item *ServiceInstance) string {
		return item.ID
	})
}

func GetCommandHandler(id, method string) *ServiceInstance {
	commandDirectoryMutex.Lock()
	defer commandDirectoryMutex.Unlock()

	ky := GetCommandKey(id, method)
	if val, ok := commandDirectory[ky]; ok {
		if len(val.Handler) == 0 {
			return nil
		}

		idx := val.robinIndex % uint(len(val.Handler))
		val.robinIndex = idx + 1
		return val.Handler[idx]
	}

	return nil
}

func RemoveCommand(id, method string) {
	commandDirectoryMutex.Lock()
	defer commandDirectoryMutex.Unlock()

	ky := GetCommandKey(id, method)
	delete(commandDirectory, ky)
}
