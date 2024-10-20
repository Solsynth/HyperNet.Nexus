package api

import (
	directory2 "git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func listExistsService(c *fiber.Ctx) error {
	services := directory2.ListServiceInstance()

	return c.JSON(lo.Map(services, func(item *directory2.ServiceInstance, index int) map[string]any {
		return map[string]any{
			"id":    item.ID,
			"type":  item.Type,
			"label": item.Label,
		}
	}))
}
