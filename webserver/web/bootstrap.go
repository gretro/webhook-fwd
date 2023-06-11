package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gretro/webhook-fwd/config"
)

var webserver *fiber.App

func BootstrapWebServer(cfg *config.AppConfig) {
	webserver = fiber.New()
	webserver.Use(
		// TODO: Implement Basic Auth to handle API Keys
		requestid.New(requestid.Config{
			Header:     "X-Request-ID",
			ContextKey: RequestIdKey,
		}),
		// TODO: Change for Zap
		logger.New(),
		recover.New(),
	)

	webserver.Get("/ready", func(ctx *fiber.Ctx) error {
		ctx.SendString("ready!")
		return nil
	})

	webserver.Get("/health", func(ctx *fiber.Ctx) error {
		ctx.SendString("healthy")
		return nil
	})

	webserver.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func WebServer() *fiber.App {
	if webserver == nil {
		panic("WebServer is not bootstrapped")
	}

	return webserver
}
