package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisConfig() RedisConfig {
	return RedisConfig{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.address"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
}
