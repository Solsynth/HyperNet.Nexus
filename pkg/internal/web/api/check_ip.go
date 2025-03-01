package api

import "github.com/gofiber/fiber/v2"

func getClientIP(c *fiber.Ctx) error {
	return c.SendString(c.IP())
}
