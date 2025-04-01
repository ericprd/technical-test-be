package redisrepo

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func (r *impl) Get(ctx context.Context, key string) (any, error) {
	val, err := r.RDB.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		return nil, fmt.Errorf("data with key %v does not exist: %v", key, err)
	case err != nil:
		return nil, fmt.Errorf("failed to get data with key %v: %v", key, err)
	case val == "":
		return nil, fmt.Errorf("data with key %v does not exist: %v", key, err)
	}

	return val, nil
}
