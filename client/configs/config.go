package config

import (
	"github.com/spf13/viper"
)

type config struct {
	PORT string // `mapstructure:"PORT"`
}

var cfg *config

func GetPort() *string {
	return &cfg.PORT
}

func init() {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath("../.")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
}
