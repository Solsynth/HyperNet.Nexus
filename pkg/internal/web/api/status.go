package api

import (
	"git.solsynth.dev/hypernet/nexus/pkg/internal/directory"
	"github.com/gofiber/fiber/v2"
)

func getServicesStatus(c *fiber.Ctx) error {
	return c.JSON(directory.GetServiceStatus())
}
