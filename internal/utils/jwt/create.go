package jwtutil

import (
	"time"

	"github.com/ericprd/technical-test/config"
	timeutil "github.com/ericprd/technical-test/internal/utils/time"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(data AuthToken) (string, error) {
	now := timeutil.UTCNow()
	claims := jwt.MapClaims{
		string(USERNAME):   data.Username,
		string(ROLE):       data.Role,
		string(SESSION_ID): data.SessionID,
		string(ID):         data.ID,
		string(EXPIRED_AT): now.Add(time.Minute * time.Duration(config.REDIS_LIFESPAN)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.SECRET_KEY))
}
