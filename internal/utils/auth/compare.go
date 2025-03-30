package authutil

import "golang.org/x/crypto/bcrypt"

func ComparePassword(storedPassword, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(inputPassword))
	return err == nil
}
