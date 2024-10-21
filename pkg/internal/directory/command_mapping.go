package directory

import (
	"context"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/kv"
	"git.solsynth.dev/hypernet/nexus/pkg/nex"
	"github.com/goccy/go-json"
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

const CommandInfoKvPrefix = "nexus.command/"

func AddCommand(id, method string, tags []string, handler *ServiceInstance) error {
	if tags == nil {
		tags = make([]string, 0)
	}

	ky := CommandInfoKvPrefix + nex.GetCommandKey(id, method)

	command := &Command{
		ID:      id,
		Method:  method,
		Tags:    tags,
		Handler: []*ServiceInstance{handler},
	}

	command.Handler = lo.UniqBy(command.Handler, func(item *ServiceInstance) string {
		return item.ID
	})

	commandJSON, err := json.Marshal(command)
	if err != nil {
		log.Printf("Error marshaling command: %v", err)
		return nil
	}

	_, err = kv.Kv.Put(context.Background(), ky, string(commandJSON))
	return err
}

func GetCommandHandler(id, method string) *ServiceInstance {
	ky := CommandInfoKvPrefix + nex.GetCommandKey(id, method)

	resp, err := kv.Kv.Get(context.Background(), ky)
	if err != nil {
		return nil
	}

	if len(resp.Kvs) == 0 {
		return nil
	}

	var command Command
	if err := json.Unmarshal(resp.Kvs[0].Value, &command); err != nil {
		return nil
	}

	if len(command.Handler) == 0 {
		return nil
	}

	idx := command.RobinIndex % uint(len(command.Handler))
	command.RobinIndex = idx + 1

	raw, err := json.Marshal(&command)
	if err == nil {
		_, _ = kv.Kv.Put(context.Background(), ky, string(raw))
	}

	return command.Handler[idx]
}

func RemoveCommand(id, method string) error {
	ky := CommandInfoKvPrefix + nex.GetCommandKey(id, method)

	_, err := kv.Kv.Delete(context.Background(), ky)
	return err
}
