package main

import (
	lflag "flag"
	los "os"
	lsignal "os/signal"

	lzap "go.uber.org/zap"
	lzapcore "go.uber.org/zap/zapcore"

	lsapiservice "github.com/cuongpiger/reforged-labs/api-service"
	lsconfig "github.com/cuongpiger/reforged-labs/configuration/api-service"
	lsversion "github.com/cuongpiger/reforged-labs/version"
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
	lzap.L().Info("Start the server")
	apiServiceConfig.Init()
	apiService, err := lsapiservice.NewAPIService(apiServiceConfig)
	if err != nil {
		lzap.L().Error("Failed to create the server", lzap.Error(err))
		los.Exit(1)
	}

	// Warm up the server
	lzap.L().Info("Warm up API service")
	apiService.WarmUp()

	// Signal stop service
	lzap.L().Info("Configure signal stop service")
	signalChan := make(chan los.Signal, 1)
	lsignal.Notify(signalChan, los.Interrupt, los.Kill)
	go func() {
		lzap.S().Info("System call to stop service: ", <-signalChan)
		if err = apiService.Stop(); err != nil {
			lzap.L().Error("Failed to stop API service", lzap.Error(err))
			panic(err)
		}
		lzap.S().Info("System call: ", <-signalChan)
	}()

	// Start the server
	lzap.L().Info("Start API service", lzap.String("version", lsversion.Get().FullyString()))
	if err = apiService.ServeHTTPService(); err != nil {
		lzap.L().Error("Failed to start API service's HTTP service", lzap.Error(err))
		panic(err)
	}
}
