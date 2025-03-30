package config

import "github.com/redis/go-redis/v9"

func newRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:            REDIS_HOST,
		Password:        REDIS_PASS,
		DB:              REDIS_DB,
		DisableIdentity: REDIS_DISABLE_ID,
	})
}
