package config

import (
	"github.com/BumwooPark/util/store"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewConfig(t *testing.T) {

	viper.SetConfigName("config")  // name of config file (without extension)
	viper.AddConfigPath("/config") // path to look for the config file in
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	require.NoError(t, err)

	store.NewDatabase()
}
