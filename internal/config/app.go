package config

import "github.com/spf13/viper"

type Application struct {
	Host     string
	HTTPPort int
}

func NewApplication() Application {
	return Application{
		Host:     viper.GetString("server.host"),
		HTTPPort: viper.GetInt("server.port"),
	}
}
