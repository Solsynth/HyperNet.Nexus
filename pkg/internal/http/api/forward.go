package api

import (
	"fmt"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
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

	url := c.OriginalURL()
	url = strings.Replace(url, "/cgi/"+ogKeyword, "", 1)
	url = *service.HttpAddr + url

	if tk, ok := c.Locals("nex_token").(string); ok {
		c.Request().Header.Set(fiber.HeaderAuthorization, fmt.Sprintf("Bearer %s", tk))
	} else {
		c.Request().Header.Del(fiber.HeaderAuthorization)
	}

	return proxy.Do(c, url, &fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		DisablePathNormalizing:   true,
		StreamResponseBody:       true,
	})
}
