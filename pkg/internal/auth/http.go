package auth

import (
	"github.com/gofiber/fiber/v2"
)

func AuthContextMiddleware(c *fiber.Ctx) error {
	atk := tokenExtract(c)
	c.Locals("nex_in_token", atk)

	if claims, err := tokenRead(atk); err == nil && claims != nil {
		c.Locals("nex_principal", claims)
		if err = userinfoFetch(c); err != nil {
			return err
		}
	} else if err != nil {
		c.Locals("nex_auth_error", err)
	}

	return c.Next()
}

func AuthMiddleware(c *fiber.Ctx) error {
	if c.Locals("nex_principal") == nil {
		err := c.Locals("nex_auth_error").(error)
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.Next()
}
