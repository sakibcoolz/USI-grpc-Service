package main

import (
	"USI-grpc-Service/config"
	"USI-grpc-Service/logger"
	"USI-grpc-Service/server"
)

func main() {
	logger := logger.GetLogger()

	// log service starting
	logger.Info("Starting service")

	// config service
	configs := config.New()

	configs.GetEnvConfig()

	server.Server(configs, logger)
}
