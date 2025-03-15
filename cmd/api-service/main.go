package main

import (
	lflag "flag"
	los "os"

	lzap "go.uber.org/zap"
	lzapcore "go.uber.org/zap/zapcore"

	lsconfig "github.com/vngcloud/reforged-labs/configuration/api-service"
)

var (
	configFilePath string
)

func setupConfigProgramParameters() {
	lflag.StringVar(&configFilePath, "config-file", "", "Define the path to the configuration file")
	lflag.Parse()
}

func setupLogging() {
	var (
		err error
		log *lzap.Logger
	)

	loggingConfiguration := lzap.NewProductionConfig()
	if los.Getenv("APP_ENV") == "development" {
		loggingConfiguration = lzap.NewDevelopmentConfig()
	}

	loggingConfiguration.EncoderConfig.TimeKey = "time"                                                        // Set the time key name
	loggingConfiguration.EncoderConfig.EncodeTime = lzapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000000") // Custom time format
	loggingConfiguration.InitialFields = map[string]interface{}{"app": "api-service"}                          // Set the initial fields
	log, err = loggingConfiguration.Build()
	if err != nil {
		los.Exit(1)
	}

	lzap.ReplaceGlobals(log)
}

func main() {
	setupConfigProgramParameters()
	setupLogging()

	// Set up the recovery mode
	defer func() {
		lzap.L().Info("Setup recovery mode")
		if err := recover(); err != nil {
			lzap.L().Error("Recover when starting the API service", lzap.Error(err.(error)))
			los.Exit(0)
		}
	}()

	// Load the configuration file
	lzap.L().Info("Load the configuration file", lzap.String("file_path", configFilePath))
	apiServiceConfig, err := lsconfig.LoadAPIServiceConfiguration(configFilePath)
	if err != nil {
		lzap.L().Error("Failed to load the configuration file", lzap.Error(err))
		los.Exit(1)
	}
	lzap.L().Info("Configuration file loaded")

	// Start the new server
	lzap.L().Info("Start the server", lzap.Any("config", apiServiceConfig))
	apiServiceConfig.Init()
}
