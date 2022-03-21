package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/to77e/sudoku-api/pkg/logging"
)

type Config struct {
	IsDebug *bool `yaml:"isDebug"`
	Listen  struct {
		Type   string `yaml:"isDebug"`
		BindIP string `yaml:"bindIp"`
		Port   string `yaml:"port"`
	} `yaml:"listen"`
	Storage StorageConfig `yaml:"storage"`
}

type StorageConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logger := logging.GetLogger()
		logger.Info("Read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yaml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			logger.Info(help)
			logger.Fatal(err)
		}
	})
	return instance
}
