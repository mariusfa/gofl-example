package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port int
}

func GetConfig() (*Config, error) {
	defaultViper := viper.New()

	defaultViper.SetConfigName("application")
	defaultViper.SetConfigType("yaml")
	defaultViper.AddConfigPath(".")

	defaultViper.AutomaticEnv()
	defaultViper.SetEnvPrefix("")

	if err := defaultViper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	err := defaultViper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
