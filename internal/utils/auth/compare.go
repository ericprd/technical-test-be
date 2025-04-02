package authutil

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(storedPassword, inputPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword)); err != nil {
		return err
	}
	return nil
}
