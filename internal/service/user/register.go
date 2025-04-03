package usersvc

import (
	"context"
	"fmt"

	userdomain "github.com/ericprd/technical-test/internal/domain/user"
	walletdomain "github.com/ericprd/technical-test/internal/domain/wallet"
	authutil "github.com/ericprd/technical-test/internal/utils/auth"
	jwtutil "github.com/ericprd/technical-test/internal/utils/jwt"
	"github.com/google/uuid"
)

func (i *impl) Register(ctx context.Context, spec userdomain.RegisterSpec) (string, error) {
	spec.ID = uuid.NewString()

	hashedPass, err := authutil.HashPassword(spec.Password)
	if err != nil {
		return "", fmt.Errorf("failed hashing password: %v", err)
	}

	spec.Password = hashedPass

	if err := i.userRepo.Register(spec); err != nil {
		return "", fmt.Errorf("register user failed: %v", err)
	}

	if err := i.walletRepo.Register(walletdomain.Wallet{
		ID:     uuid.NewString(),
		UserID: spec.ID,
	}); err != nil {
		return "", fmt.Errorf("failed to create user's wallet: %v", err)
	}

	token, err := jwtutil.CreateJWT(jwtutil.AuthToken{
		Username:  spec.Username,
		SessionID: spec.ID,
		Role:      "user",
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	key := fmt.Sprintf("user-%s", spec.ID)

	if err := i.rdb.Set(ctx, token, key); err != nil {
		return "", fmt.Errorf("failed to cache token: %v", err)
	}
	return token, nil
}
