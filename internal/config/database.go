package config

import "github.com/spf13/viper"

type Database struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

func DefaultDB() *Database {
	return &Database{
		Name:     viper.GetString("database.name"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
	}
}
