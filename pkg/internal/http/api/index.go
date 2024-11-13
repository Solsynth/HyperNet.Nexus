package api

import (
	pkg "git.solsynth.dev/hypernet/nexus/pkg/internal"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/auth"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/http/ws"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func MapAPIs(app *fiber.App) {
	// Some built-in public-accessible APIs
	wellKnown := app.Group("/.well-known").Name("Well Known")
	{
		wellKnown.Get("/", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"api_level": pkg.ApiLevel,
				"version":   pkg.AppVersion,
				"status":    true,
			})
		})
		wellKnown.Get("/check-ip", getClientIP)
		wellKnown.Get("/directory/services", listExistsService)
	}

	// Common websocket gateway
	app.Get("/ws", auth.ValidatorMiddleware, websocket.New(ws.Listen))

	app.All("/inv/:command", invokeCommand)
	app.All("/cgi/:service/*", forwardService)
}
