package config

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {

	os.Setenv("MODE", "debug")
	t.Log(os.Getenv("MODE"))
	t.Log(viper.Get("MODE"))

	viper.SetConfigName("config")  // name of config file (without extension)
	viper.AddConfigPath("/config") // path to look for the config file in
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	require.NoError(t, err)
}
