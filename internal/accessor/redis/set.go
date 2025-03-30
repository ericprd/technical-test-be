package redisrepo

import (
	"context"
	"time"

	"github.com/ericprd/technical-test/config"
	"github.com/redis/go-redis/v9"
)

func (r *impl) Set(ctx context.Context, data interface{}, key string) error {
	_, err := r.RDB.Get(ctx, key).Result()

	if err != redis.Nil {
		if err := r.RDB.Del(ctx, key).Err(); err != nil {
			return err
		}
	}

	return r.RDB.Set(ctx, key, data, time.Millisecond*time.Duration(config.REDIS_LIFESPAN)).Err()
}
