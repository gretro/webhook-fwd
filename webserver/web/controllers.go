package web

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gretro/webhook-fwd/webserver/web/controllers"
)

type ControllerRegistry struct {
	channel *controllers.ChannelController
}

var registry *ControllerRegistry

func BootstrapControllers(api *fiber.App) {
	registry = &ControllerRegistry{
		channel: controllers.NewChannelController(),
	}

	registry.channel.Register(api)
}
