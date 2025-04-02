package usersvc

import (
	"context"
	"fmt"

	jwtutil "github.com/ericprd/technical-test/internal/utils/jwt"
)

func (i *impl) Logout(ctx context.Context, token string) error {
	claims, err := jwtutil.VerifyJWT(token)
	if err != nil {
		return fmt.Errorf("failed to verify token: %v", err)
	}

	id, ok := claims[string(jwtutil.SESSION_ID)]
	if !ok {
		return fmt.Errorf("failed to get session id")
	}

	key := fmt.Sprintf("user-%s", id)

	if err := i.rdb.Delete(ctx, key); err != nil {
		return fmt.Errorf("failed to delete stored token: %v", err)
	}

	return nil
}
