package redisrepo

import "context"

func (r *impl) Delete(ctx context.Context, key string) error {
	return r.RDB.Del(ctx, key).Err()
}
