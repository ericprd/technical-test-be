package authutil

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	// Hashing the password with a cost of 14 (default)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
