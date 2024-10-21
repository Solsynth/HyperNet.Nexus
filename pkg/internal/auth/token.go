package auth

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var JReader *sec.JwtReader

func SoftAuthMiddleware(c *fiber.Ctx) error {
	atk := tokenExtract(c)
	c.Locals("nex_token", atk)

	if claims, err := tokenRead(atk); err == nil && claims != nil {
		c.Locals("nex_principal", claims)
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

func tokenExtract(c *fiber.Ctx) string {
	var atk string
	if cookie := c.Cookies(sec.CookieAccessToken); len(cookie) > 0 {
		atk = cookie
	}
	if header := c.Get(fiber.HeaderAuthorization); len(header) > 0 {
		tk := strings.Replace(header, "Bearer", "", 1)
		atk = strings.TrimSpace(tk)
	}
	if tk := c.Query("tk"); len(tk) > 0 {
		atk = strings.TrimSpace(tk)
	}
	return atk
}

func tokenRead(in string) (*sec.JwtClaims, error) {
	if JReader == nil {
		return nil, nil
	}

	claims, err := sec.ReadJwt[sec.JwtClaims](JReader, in)
	return &claims, err
}
