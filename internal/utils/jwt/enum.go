package jwtutil

type JWTClaim string

const (
	USERNAME   JWTClaim = "username"
	ROLE       JWTClaim = "role"
	SESSION_ID JWTClaim = "sid"
	ID         JWTClaim = "id"
	EXPIRED_AT JWTClaim = "exp"
)
