package commands

import (
	"errors"

	"github.com/gretro/webhook-fwd/cli/client"
	apperrors "github.com/gretro/webhook-fwd/errors"
	"github.com/gretro/webhook-fwd/libs"
	"github.com/gretro/webhook-fwd/utils"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func ConnectCmd(ctx *cli.Context) error {
	channelName := ctx.String(ChannelFlag)
	if channelName == "" {
		channelName = utils.GenerateRandomName()
	}

	server := ctx.String(ServerFlag)
	if server == "" {
		server = DefaultServer
	}

	retries := ctx.Int(RetryFlag)

	apiClient := client.NewClient(client.ClientOptions{
		Authorization: nil,
		ServerUrl:     server,
		Agent:         GetUserAgent(ctx),
		RetryCount:    retries,
		RetryDelay:    DefaultRetryDuration,
	})

	channelRef := apiClient.Channel(channelName)

	_, err := channelRef.Create(ctx.Context)
	if err != nil && !errors.As(err, &apperrors.AlreadyExistsError{}) {
		libs.Logger().Error("Could not create the channel", zap.Error(err), zap.String("channel", channelName))
		return err
	}

	libs.Logger().Info("Channel created", zap.String("channel", channelName))

	// TODO: Implement connect logic

	return nil
}
