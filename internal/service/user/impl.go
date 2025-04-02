package usersvc

import (
	"context"

	redisrepo "github.com/ericprd/technical-test/internal/accessor/redis"
	userrepo "github.com/ericprd/technical-test/internal/accessor/user"
	walletrepo "github.com/ericprd/technical-test/internal/accessor/wallet"
	userdomain "github.com/ericprd/technical-test/internal/domain/user"
)

type impl struct {
	rdb        redisrepo.Redis
	userRepo   userrepo.User
	walletRepo walletrepo.Wallet
}

type User interface {
	Register(ctx context.Context, spec userdomain.Request) (string, error)
}

func New(
	rdb redisrepo.Redis,
	userRepo userrepo.User,
	walletRepo walletrepo.Wallet,
) User {
	return &impl{rdb, userRepo, walletRepo}
}
