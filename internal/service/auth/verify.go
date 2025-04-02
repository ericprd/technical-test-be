package authsvc

import (
	"fmt"
	"net/http"
	"time"

	authutil "github.com/ericprd/technical-test/internal/utils/auth"
	jwtutil "github.com/ericprd/technical-test/internal/utils/jwt"
)

func (i *impl) Verify(r *http.Request) error {
	ctx := r.Context()

	token, err := authutil.GetTokenHeader(r, "Authorization")
	if err != nil {
		return fmt.Errorf("failed to get token: %v", err)
	}

	claims, err := jwtutil.VerifyJWT(token)
	if err != nil {
		return fmt.Errorf("failed to get claims: %v", err)
	}

	id, ok := claims[string(jwtutil.SESSION_ID)]
	if !ok {
		return fmt.Errorf("authorization is required")
	}

	key := fmt.Sprintf("user-%s", id)

	storedToken, err := i.rdb.Get(ctx, key).Result()
	if err != nil {
		return fmt.Errorf("failed to get stored token: %v", err)
	}

	if storedToken != token {
		return fmt.Errorf("invalid token")
	}

	expTime := claims[string(jwtutil.EXPIRED_AT)].(float64)

	if time.Unix(int64(expTime), 0).Before(time.Now()) {
		return fmt.Errorf("token has expired")
	}

	return nil
}
