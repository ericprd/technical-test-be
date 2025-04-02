package authsvc

import (
	"net/http"

	"github.com/redis/go-redis/v9"
)

type impl struct {
	rdb *redis.Client
}

type Auth interface {
	Verify(r *http.Request) error
}

func New(rdb *redis.Client) Auth {
	return &impl{rdb}
}
