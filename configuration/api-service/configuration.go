package api_service

import (
	los "os"

	lzap "go.uber.org/zap"
	lyaml "gopkg.in/yaml.v3"
)

type APIServiceConfiguration struct {
	APIService struct {
		Port int    `yaml:"port,omitempty"`
		Host string `yaml:"host,omitempty"`

		Database struct {
			URI string `yaml:"uri,omitempty"`
		} `yaml:"database,omitempty"`

		WorkerPool struct {
			Amount int `yaml:"amount,omitempty"`
			Buffer int `yaml:"buffer,omitempty"`
		} `yaml:"worker_pool,omitempty"`
	} `yaml:"api_service,omitempty"`
}

func (s *APIServiceConfiguration) Init() {
	if s.APIService.Port == 0 {
		s.APIService.Port = 8000
	}

	if s.APIService.Host == "" {
		s.APIService.Host = "0.0.0.0"
	}
}

func LoadAPIServiceConfiguration(pfilePath string) (*APIServiceConfiguration, error) {
	if len(pfilePath) < 1 {
		pfilePath = los.Getenv("API_SERVICE_CONFIG_FILE")
	}

	lzap.L().Info("Load configuration file", lzap.String("file_path", pfilePath))
	configBytes, err := los.ReadFile(pfilePath)
	if err != nil {
		lzap.L().Error("Failed to load configuration file", lzap.Error(err))
		return nil, err
	}

	configBytes = []byte(los.ExpandEnv(string(configBytes)))
	cfg := new(APIServiceConfiguration)
	err = lyaml.Unmarshal(configBytes, cfg)
	if err != nil {
		lzap.L().Error("Failed to parse configuration file", lzap.Error(err))
		return nil, err
	}

	lzap.L().Info("Configuration file loaded successfully")
	return cfg, nil
}
