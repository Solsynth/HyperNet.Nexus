package api

import (
	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"strings"
)

func forwardService(c *fiber.Ctx) error {
	serviceType := c.Params("service")
	ogKeyword := serviceType

	aliasingMap := viper.GetStringMapString("services.aliases")
	if val, ok := aliasingMap[serviceType]; ok {
		serviceType = val
	}

	service := directory.GetServiceInstanceByType(serviceType)

	if service == nil || service.HttpAddr == nil {
		return fiber.NewError(fiber.StatusNotFound, "service not found")
	}

	ogUrl := c.Request().URI().String()
	url := c.OriginalURL()
	url = strings.Replace(url, "/cgi/"+ogKeyword, "", 1)
	url = *service.HttpAddr + url

	log.Debug().
		Str("from", ogUrl).
		Str("to", url).
		Str("service", serviceType).
		Str("id", service.ID).
		Msg("Forwarding request for service...")

	return proxy.Do(c, url)

}
