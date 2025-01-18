package cache

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"time"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) Cache {
	return &Redis{
		client: client,
	}
}

func (r *Redis) Ping() error {
	return r.client.Ping().Err()
}

func (r *Redis) Get(key string) (string, error) {
	resStr, err := r.client.Get(key).Result()

	if err != nil {
		return "", err
	}

	return resStr, nil
}

func (r *Redis) Set(key string, val interface{}, exp time.Duration) error {
	bytesValue, err := json.Marshal(val)

	if err != nil {
		return err
	}

	return r.client.Set(key, bytesValue, exp).Err()
}
