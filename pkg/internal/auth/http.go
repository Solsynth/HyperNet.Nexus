package auth

import "github.com/gofiber/fiber/v2"

func SoftAuthMiddleware(c *fiber.Ctx) error {
	atk := tokenExtract(c)
	c.Locals("nex_token", atk)

	if claims, err := tokenRead(atk); err == nil && claims != nil {
		c.Locals("nex_principal", claims)
		// TODO fetch user info
	} else if err != nil {
		c.Locals("nex_auth_error", err)
	}

	return c.Next()
}

func HardAuthMiddleware(c *fiber.Ctx) error {
	if c.Locals("nex_principal") == nil {
		err := c.Locals("nex_auth_error").(error)
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.Next()
}
