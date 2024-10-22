package sec

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

// ContextMiddleware provide a middleware to receive the userinfo from the nexus.
// It only works on the client-side of nexus.
// It will NOT validate the auth status if you need to validate the status of current authorization, refer to ValidatorMiddleware.
// To get the userinfo, call `c.Locals('nex_user').(sec.UserInfo)`
// Make sure you got the right public key, otherwise the auth will fail.
func ContextMiddleware(tkReader *InternalTokenReader) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get(fiber.HeaderAuthorization)
		token = strings.TrimSpace(strings.Replace(token, "Bearer ", "", 1))
		if len(token) == 0 {
			return fiber.NewError(fiber.StatusUnauthorized, "no authorization token is provided")
		}

		data, err := tkReader.ReadUserInfoJwt(token)
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, err.Error())
		}
		c.Locals("nex_user", data)

		return c.Next()
	}
}

// ValidatorMiddleware will ensure request is authenticated
// Make sure call this middleware after ContextMiddleware
func ValidatorMiddleware(c *fiber.Ctx) error {
	if c.Locals("nex_user") == nil {
		return fiber.NewError(fiber.StatusUnauthorized, "unauthorized")
	}

	return c.Next()
}