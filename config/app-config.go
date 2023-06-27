package config

import "go.uber.org/zap"

const (
	DevelopmentAppEnv = "development"
	ProductionAppEnv  = "production"
)

type AppConfig struct {
	AppEnvironment string
	HttpPort       int
	PublicHost     string
	LogLevel       string
}

var appConfig *AppConfig

func BoostrapAppConfiguration() *AppConfig {
	// TODO: Read from environment variables
	appConfig = &AppConfig{
		AppEnvironment: DevelopmentAppEnv,
		HttpPort:       5333,
		PublicHost:     "http://localhost:5333",
		LogLevel:       zap.InfoLevel.String(),
	}

	return appConfig
}

func AppConfiguration() *AppConfig {
	if appConfig == nil {
		panic("AppConfiguration is not bootstrapped")
	}

	return appConfig
}
