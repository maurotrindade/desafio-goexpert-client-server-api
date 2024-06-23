package config

import (
	"github.com/spf13/viper"
)

type config struct {
	PORT string // `mapstructure:"PORT"`

	DB_DRIVER string
	DB_HOST   string
	DB_PORT   string
	DB_USER   string
	DB_PSW    string
	DB_NAME   string
}

var cfg *config

func GetPort() *string {
	return &cfg.PORT
}

func GetDbDriver() *string {
	return &cfg.DB_DRIVER
}

func GetDbHost() *string {
	return &cfg.DB_HOST
}

func GetDbPort() *string {
	return &cfg.DB_PORT
}

func GetDbUser() *string {
	return &cfg.DB_USER
}

func GetDbPsw() *string {
	return &cfg.DB_PSW
}

func GetDbName() *string {
	return &cfg.DB_NAME
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
