package api

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func MapAPIs(app *fiber.App) {
	// Some built-in public-accessible APIs
	wellKnown := app.Group("/.well-known").Name("Well Known")
	{
		wellKnown.Get("/", func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusOK)
		})
		wellKnown.Get("/check-ip", getClientIP)
		wellKnown.Get("/directory/services", listExistsService)
	}

	// Common websocket gateway
	app.Use(func(c *fiber.Ctx) error {
		/*if err := exts.EnsureAuthenticated(c); err != nil {
			return err
		}*/
		return c.Next()
	}).Get("/ws", websocket.New(listenWebsocket))

	app.All("/cgi/:service/*", forwardServiceRequest)
}
