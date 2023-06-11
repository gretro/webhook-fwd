package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

const (
	UserAgent = "webhook-fwd-cli"
)

func GetUserAgent(ctx *cli.Context) string {
	return fmt.Sprintf("%s/%s", UserAgent, ctx.App.Version)
}
