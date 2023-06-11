package libs

import (
	"fmt"

	"github.com/gretro/webhook-fwd/src/config"
	"go.uber.org/zap"
)

var logger *zap.Logger

type CLILoggerOptions struct {
	Quiet   bool
	Verbose bool
	Debug   bool
}

func BootstrapWebLogger(cfg *config.AppConfig) *zap.Logger {
	atomicLevel, err := zap.ParseAtomicLevel(cfg.LogLevel)
	var level = zap.InfoLevel
	if err == nil {
		level = atomicLevel.Level()
	}

	if cfg.AppEnvironment == config.DevelopmentAppEnv {
		logger, err = zap.NewDevelopment(
			zap.IncreaseLevel(level),
		)
	} else {
		logger, err = zap.NewProduction(
			zap.IncreaseLevel(level),
		)
	}

	if err != nil {
		panic(fmt.Errorf("could not initialize the logger: %w", err))
	}

	return logger
}

func BootstrapCLILogger(options CLILoggerOptions) *zap.Logger {
	if options.Quiet {
		logger = zap.NewNop()
	} else {
		zapOptions := make([]zap.Option, 0)

		if !options.Debug {
			logLevel := zap.WarnLevel

			if options.Verbose {
				logLevel = zap.InfoLevel
			}

			zapOptions = append(zapOptions, zap.IncreaseLevel(logLevel))
		}

		if options.Debug {
			zapOptions = append(zapOptions, zap.AddStacktrace(zap.ErrorLevel))
		} else {
			zapOptions = append(zapOptions, zap.AddStacktrace(zap.FatalLevel))
		}

		var err error
		logger, err = zap.NewDevelopment(zapOptions...)
		if err != nil {
			panic(fmt.Errorf("could not initialize the logger: %w", err))
		}
	}

	return logger
}

func Logger() *zap.Logger {
	if logger == nil {
		panic("Logger is not initialized")
	}

	return logger
}
