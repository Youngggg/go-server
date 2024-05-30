package mw

import (
	"time"

	"apple/common/cfg"

	"github.com/redis/go-redis/v9"
)

func initRedis(config cfg.RedisConfig) *redis.Client {
	op := &redis.Options{
		DB:              config.Db,
		Addr:            config.Address,
		Password:        config.Password,
		DialTimeout:     5 * time.Second,
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    5 * time.Second,
		PoolSize:        100,
		MaxRetries:      3,
		MinRetryBackoff: -1,
		MaxRetryBackoff: -1,
	}

	return redis.NewClient(op)
}
