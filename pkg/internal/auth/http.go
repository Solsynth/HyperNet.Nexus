package auth

import (
	"github.com/gofiber/fiber/v2"
)

func ContextMiddleware(c *fiber.Ctx) error {
	atk := tokenExtract(c)
	if len(atk) == 0 {
		return c.Next()
	}

	c.Locals("nex_in_token", atk)

	if claims, err := tokenRead(atk); err == nil && claims != nil {
		c.Locals("nex_principal", claims)
		if err = userinfoFetch(c); err != nil {
			return err
		}
	} else if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.Next()
}

func ValidatorMiddleware(c *fiber.Ctx) error {
	if c.Locals("nex_principal") == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	return c.Next()
}
