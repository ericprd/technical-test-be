package redisrepo

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type impl struct {
	RDB *redis.Client
}

type Redis interface {
	Set(ctx context.Context, data interface{}, key string) error
	Get(ctx context.Context, key string) (interface{}, error)
	Delete(ctx context.Context, key string) error
}

func New(
	rdb *redis.Client,
) Redis {
	return &impl{rdb}
}
