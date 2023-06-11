package web

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gretro/webhook-fwd/config"
)

var apiServer *fiber.App

func BootstrapWebServer(cfg *config.AppConfig) {
	apiServer = fiber.New()
	apiServer.Use(
		// TODO: Implement Basic Auth to handle API Keys
		requestid.New(requestid.Config{
			Header:     "X-Request-ID",
			ContextKey: RequestIdKey,
		}),
		// TODO: Change for Zap
		logger.New(),
		recover.New(),
	)

	apiServer.Get("/ready", func(ctx *fiber.Ctx) error {
		ctx.SendString("ready!")
		return nil
	})

	apiServer.Get("/health", func(ctx *fiber.Ctx) error {
		ctx.SendString("healthy")
		return nil
	})

	BootstrapControllers(apiServer)

	apiServer.Listen(fmt.Sprintf(":%d", cfg.HttpPort))
}

func ApiServer() *fiber.App {
	if apiServer == nil {
		panic("WebServer is not bootstrapped")
	}

	return apiServer
}
