package jwtutil

import (
	"fmt"

	"github.com/ericprd/technical-test/config"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyJWT(s string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(s, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
