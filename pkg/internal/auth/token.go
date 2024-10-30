package auth

import (
	"git.solsynth.dev/hypernet/nexus/pkg/nex/sec"
	"github.com/gofiber/fiber/v2"
	"strings"
)

var JReader *sec.JwtReader

var IReader *sec.InternalTokenReader
var IWriter *sec.InternalTokenWriter

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

	claims, err := sec.ReadJwt[*sec.JwtClaims](JReader, in, &sec.JwtClaims{})
	return claims, err
}
