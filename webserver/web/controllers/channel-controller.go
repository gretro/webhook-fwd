package controllers

import "github.com/gofiber/fiber/v2"

type ChannelController struct {
}

func NewChannelController() *ChannelController {
	return &ChannelController{}
}

func (controller *ChannelController) Register(apiServer *fiber.App) {
	group := apiServer.Group("/channels")

	group.Get("/:channelName", controller.GetChannel)
}

func (controller *ChannelController) GetChannel(ctx *fiber.Ctx) error {
	// channelName := ctx.Params("channelName")

	ctx.SendString("ok")
	return nil
}
