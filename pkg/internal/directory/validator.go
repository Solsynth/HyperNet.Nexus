package directory

import "github.com/rs/zerolog/log"

func ValidateServices() {
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
			_ = RemoveServiceInstance(service.ID)
			log.Warn().Err(err).Str("id", service.ID).Str("addr", service.GrpcAddr).Msg("Unable connect to service, dropped...")
			continue
		}

		successCount++
	}

	log.Info().
		Int("success", successCount).
		Int("failed", len(services)-successCount).
		Int("total", len(services)).
		Msg("Service validation completed.")
}
