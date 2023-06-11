package main

import (
	"fmt"
	"os"

	"github.com/gretro/webhook-fwd/src/commands"
	"github.com/gretro/webhook-fwd/src/libs"
	"github.com/urfave/cli/v2"
)

const (
	Version = "1.0.0"

	LoggingCategory = "Logging"
	ConfigCategory  = "Configuration"

	ServerUrlEnv = "WEBHOOK_FWD_SERVER"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}()

	app := &cli.App{
		Name:                 "webhook-fwd",
		Usage:                "Forward your webhooks",
		Version:              Version,
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:     commands.QuietFlag,
				Aliases:  []string{"q"},
				Category: LoggingCategory,
				Usage:    "emits no logs",
				Value:    false,
			},
			&cli.BoolFlag{
				Name:     commands.VerboseFlag,
				Category: LoggingCategory,
				Usage:    "emits verbose logs",
				Value:    false,
			},
			&cli.BoolFlag{
				Name:     commands.DebugFlag,
				Category: LoggingCategory,
				Usage:    "enables debug mode",
				Value:    false,
			},
			&cli.StringFlag{
				Name:        commands.ServerFlag,
				Category:    ConfigCategory,
				Usage:       "sets the server to target",
				Value:       commands.DefaultServer,
				DefaultText: commands.DefaultServer,
				EnvVars:     []string{ServerUrlEnv},
			},
			&cli.IntFlag{
				Name:        commands.RetryFlag,
				Category:    ConfigCategory,
				Usage:       "sets the amount of retries for each request",
				Value:       commands.DefaultRetry,
				DefaultText: fmt.Sprintf("%d", commands.DefaultRetry),
			},
		},
		Before: bootstrapCLI,
		After:  after,
		Commands: []*cli.Command{
			&cli.Command{
				Name:      "connect",
				Usage:     "Connects to a Webhook FWD server and starts forwarding requests",
				ArgsUsage: " ",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    commands.ChannelFlag,
						Aliases: []string{"c"},
						Usage:   "sets the channel to use. If the channel does not exist, it will be created. If not set, a random name will be picked",
					},
				},
				Action: commands.ConnectCmd,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}

func bootstrapCLI(ctx *cli.Context) error {
	libs.BootstrapCLILogger(libs.CLILoggerOptions{
		Verbose: ctx.Bool(commands.VerboseFlag),
		Debug:   ctx.Bool(commands.DebugFlag),
		Quiet:   ctx.Bool(commands.QuietFlag),
	})

	return nil
}

func after(ctx *cli.Context) error {
	err := libs.Logger().Sync()
	return err
}
