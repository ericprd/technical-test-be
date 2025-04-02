package usersvc

import (
	"context"
	"fmt"

	userdomain "github.com/ericprd/technical-test/internal/domain/user"
	authutil "github.com/ericprd/technical-test/internal/utils/auth"
	jwtutil "github.com/ericprd/technical-test/internal/utils/jwt"
)

func (i *impl) Login(ctx context.Context, spec userdomain.LoginSpec) (string, error) {
	user, err := i.userRepo.Login(spec.Username)
	if err != nil {
		return "", fmt.Errorf("login failed: %v", err)
	}

	if err := authutil.ComparePassword(user.Password, spec.Password); err != nil {
		return "", fmt.Errorf("wrong password: %v", err)
	}

	token, err := jwtutil.CreateJWT(jwtutil.AuthToken{
		Username:  spec.Username,
		SessionID: user.ID,
		Role:      "user",
	})

	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	key := fmt.Sprintf("user-%s", user.ID)

	if err := i.rdb.Set(ctx, token, key); err != nil {
		return "", fmt.Errorf("failed to cache token: %v", err)
	}

	return token, nil
}
