package web

import (
	"time"

	"git.solsynth.dev/hypernet/nexus/pkg/internal/auth"
	"git.solsynth.dev/hypernet/nexus/pkg/internal/web/api"
	"github.com/goccy/go-json"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type WebApp struct {
	app *fiber.App
}

func NewServer() *WebApp {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
		EnableIPValidation:    true,
		ServerHeader:          "Hypernet.Nexus",
		AppName:               "Hypernet.Nexus",
		ProxyHeader:           fiber.HeaderXForwardedFor,
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		BodyLimit:             512 * 1024 * 1024 * 1024, // 512 TiB
		EnablePrintRoutes:     viper.GetBool("debug.print_routes"),
	})

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &log.Logger,
	}))

	app.Use(idempotency.New())
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowOriginsFunc: func(origin string) bool {
			return true
		},
	}))

	app.Use(auth.ContextMiddleware)
	app.Use(limiter.New(limiter.Config{
		Max:               viper.GetInt("rate_limit"),
		Expiration:        60 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	api.MapControllers(app)

	return &WebApp{app}
}

func (v *WebApp) Listen() {
	if err := v.app.Listen(viper.GetString("bind")); err != nil {
		log.Fatal().Err(err).Msg("An error occurred when starting server...")
	}
}
