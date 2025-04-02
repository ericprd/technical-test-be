package jwtutil

import (
	"fmt"
	"testing"

	"github.com/ericprd/technical-test/config"
	"github.com/golang-jwt/jwt/v5"
)

func TestJWTGenerationAndVerification(t *testing.T) {
	config.SECRET_KEY = "HelloWorld"

	tests := []struct {
		name   string
		params AuthToken
		exec   func(args AuthToken) error
	}{
		{
			name:   "Generate and Verify Valid JWT",
			params: AuthToken{Username: "user1", SessionID: "session123", Role: "admin"},
			exec: func(authToken AuthToken) error {
				config.REDIS_LIFESPAN = 6000
				tokenString, err := CreateJWT(authToken)
				if err != nil {
					return fmt.Errorf("CreateJWT returned error: %v", err)
				}

				claims, err := VerifyJWT(tokenString)

				if err != nil {
					return fmt.Errorf("VerifyJWT returned error: %v", err)
				}

				if claims[string(USERNAME)] != authToken.Username {
					return fmt.Errorf("expected username %v, got %v", authToken.Username, claims[string(USERNAME)])
				}
				if claims[string(SESSION_ID)] != authToken.SessionID {
					return fmt.Errorf("expected sessionID %v, got %v", authToken.SessionID, claims[string(SESSION_ID)])
				}
				if claims[string(ROLE)] != authToken.Role {
					return fmt.Errorf("expected role %v, got %v", authToken.Role, claims[string(ROLE)])
				}

				return nil
			},
		},
		{
			name:   "Verify Expired JWT",
			params: AuthToken{Username: "user1", SessionID: "session123", Role: "admin"},
			exec: func(authToken AuthToken) error {
				config.REDIS_LIFESPAN = 0
				tokenString, err := CreateJWT(authToken)
				if err != nil {
					return fmt.Errorf("CreateJWT returned error: %v", err)
				}

				_, err = VerifyJWT(tokenString)
				if err == nil {
					return fmt.Errorf("expected error due to expired token, got nil")
				}

				return nil
			},
		},
		{
			name:   "Verify Invalid JWT Format",
			params: AuthToken{Username: "user1", SessionID: "session123", Role: "admin"},
			exec: func(authToken AuthToken) error {
				tokenString := "invalidtoken"

				_, err := VerifyJWT(tokenString)
				if err == nil {
					return fmt.Errorf("expected error due to invalid token format, got nil")
				}

				return nil
			},
		},
		{
			name:   "Verify JWT with Incorrect Secret",
			params: AuthToken{Username: "user1", SessionID: "session123", Role: "admin"},
			exec: func(authToken AuthToken) error {
				tokenString, err := CreateJWT(authToken)
				if err != nil {
					return fmt.Errorf("CreateJWT returned error: %v", err)
				}

				incorrectSigningKey := []byte("incorrect_secret")
				_, err = jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
					return incorrectSigningKey, nil
				})

				if err == nil {
					return fmt.Errorf("expected error due to incorrect secret key, got nil")
				}

				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.exec(tt.params)
			if err != nil {
				t.Errorf("test '%s' failed: %v", tt.name, err)
			}
		})
	}
}
