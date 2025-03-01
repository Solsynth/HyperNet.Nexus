package api

import (
	"git.solsynth.dev/hypernet/nexus/pkg/internal/watchtower"
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/gofiber/fiber/v2"
)

func wtRunDbMaintenance(c *fiber.Ctx) error {
	if err := sec.EnsureGrantedPerm(c, "AdminOperateWatchTower", true); err != nil {
		return err
	}
	go watchtower.RunDbMaintenance()
	return c.SendStatus(fiber.StatusOK)
}
