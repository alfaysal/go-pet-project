package conn

import (
	"github.com/alfaysal/go-pet-project/internal/cache"
	"github.com/alfaysal/go-pet-project/internal/config"
	"github.com/go-redis/redis"
)

var defaultCache cache.Cache
var redisClient *redis.Client

func ConnectDefaultRedis() error {
	cfg := config.NewRedisConfig()
	err := ConnectRedis(cfg)

	return err
}

func ConnectRedis(redisConfig config.RedisConfig) error {
	client := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	defaultCache = cache.NewRedis(client)
	redisClient = client

	return client.Ping().Err()
}

func DefaultCache() cache.Cache {
	return defaultCache
}
