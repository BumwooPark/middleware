package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func NewConfig(name string) {
	viper.SetConfigName(name) // name of config file (without extension)
	viper.AddConfigPath("$HOME/config")
	viper.AddConfigPath("/config") // path to look for the config file in
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
}
