package directory

import (
	"sync"

	"github.com/rs/zerolog/log"
)

var (
	statusOfServices = make(map[string]bool)
	statusLock       sync.Mutex
)

func GetServiceStatus() map[string]bool {
	out := make(map[string]bool)
	for k, v := range statusOfServices {
		out[k] = v
	}

	services := ListServiceInstance()
	for _, service := range services {
		if _, ok := out[service.Type]; !ok {
			out[service.Type] = false
		}
	}

	return out
}

func SetServiceStatus(t string, status bool) {
	statusLock.Lock()
	defer statusLock.Unlock()
	statusOfServices[t] = status
}

func ValidateServices() {
	statusLock.Lock()
	defer statusLock.Unlock()

	services := ListServiceInstance()
	if len(services) == 0 {
		return
	}

	checklist := make(map[string]bool)
	successCount := 0
	log.Info().Int("count", len(services)).Msg("Validating services...")
	for _, service := range services {
		if _, ok := checklist[service.GrpcAddr]; ok {
			_ = RemoveServiceInstance(service.ID)
			log.Warn().Str("id", service.ID).Str("addr", service.GrpcAddr).Msg("Duplicated service address, dropped...")
			continue
		}
		// Directly use the connect method to skip cache
		if _, err := ConnectService(service); err != nil {
			statusOfServices[service.Type] = false
			_ = RemoveServiceInstance(service.ID)
			log.Warn().Err(err).Str("id", service.ID).Str("addr", service.GrpcAddr).Msg("Unable connect to service, dropped...")
			continue
		} else {
			statusOfServices[service.Type] = true
		}

		successCount++
	}

	log.Info().
		Int("success", successCount).
		Int("failed", len(services)-successCount).
		Int("total", len(services)).
		Msg("Service validation completed.")
}
