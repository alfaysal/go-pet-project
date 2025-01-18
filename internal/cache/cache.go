package cache

import (
	"time"
)

type Cache interface {
	Ping() error
	Set(key string, value interface{}, exp time.Duration) error
	Get(key string) (string, error)
}
