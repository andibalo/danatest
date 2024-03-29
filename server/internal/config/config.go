package config

import (
	"os"
	logger2 "streaming/internal/logger"

	"go.uber.org/zap"
)

const envVarEnvironment = "ENV"

func InitConfig() *AppConfig {
	logger := logger2.NewMainLoggerSingleton()

	pgConf := LoadDBConfig()

	return &AppConfig{
		logger:      logger,
		environment: os.Getenv(envVarEnvironment),
		pgConfig:    pgConf,
	}
}

type AppConfig struct {
	logger      *zap.Logger
	environment string

	pgConfig *DBConfig
}

func (a *AppConfig) Logger() *zap.Logger {
	return a.logger
}

func (a *AppConfig) ServerAddress() string {
	return "8080"
}

func (a *AppConfig) PgConfig() *DBConfig {
	return a.pgConfig
}
