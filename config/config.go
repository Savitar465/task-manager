package config

import (
	"github.com/spf13/viper"
)

func InitViper() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
