package api

import (
	"git.solsynth.dev/hypernet/nexus/pkg/internal/captcha"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/web/exts"
	"github.com/gofiber/fiber/v2"
)

func renderCaptcha(c *fiber.Ctx) error {
	return c.Render("captcha", captcha.GetTemplateData())
}

func validateCaptcha(c *fiber.Ctx) error {
	var body struct {
		CaptchaToken string `json:"captcha_tk"`
	}
	if err := exts.BindAndValidate(c, &body); err != nil {
		return err
	}

	if !captcha.Validate(body.CaptchaToken, c.IP()) {
		return c.SendStatus(fiber.StatusNotAcceptable)
	}
	return c.SendStatus(fiber.StatusOK)
}
