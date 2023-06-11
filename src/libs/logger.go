package libs

import (
	"fmt"

	"go.uber.org/zap"
)

var logger *zap.Logger

type CLILoggerOptions struct {
	Quiet   bool
	Verbose bool
	Debug   bool
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
