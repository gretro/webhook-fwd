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
		HttpPort:       25333,
		PublicHost:     "http://localhost:25333",
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
