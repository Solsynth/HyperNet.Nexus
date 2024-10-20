package directory

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
	"strings"
	"sync"
)

// In commands, we use the map and the mutex because it is usually read and only sometimes write
var commandDirectory = make(map[string]*Command)
var commandDirectoryMutex sync.Mutex

func AddCommand(id, method string, tags []string, handler *ServiceInstance) {
	commandDirectoryMutex.Lock()
	defer commandDirectoryMutex.Unlock()

	if tags == nil {
		tags = make([]string, 0)
	}

	ky := nex.GetCommandKey(id, method)
	if _, ok := commandDirectory[ky]; !ok {
		commandDirectory[ky] = &Command{
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

	log.Info().Str("id", id).Str("method", method).Str("tags", strings.Join(tags, ",")).Msg("New command registered")
}

func GetCommandHandler(id, method string) *ServiceInstance {
	commandDirectoryMutex.Lock()
	defer commandDirectoryMutex.Unlock()

	ky := nex.GetCommandKey(id, method)
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

	ky := nex.GetCommandKey(id, method)
	delete(commandDirectory, ky)
}
