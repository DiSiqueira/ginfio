package config

import (
	"github.com/spf13/viper"
)

func Read() error {
	viper.SetConfigName("ginfio.yaml")
	viper.AddConfigPath("/etc/ginfio/")
	viper.AddConfigPath("$HOME/.ginfio")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	return viper.ReadInConfig()
	// @TODO Watch config changes
}
