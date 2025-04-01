package authutil

import (
	"fmt"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// TestHashAndComparePassword runs multiple tests on HashPassword and ComparePassword
func TestHashAndComparePassword(t *testing.T) {
	tests := []struct {
		name   string
		params []string
		exec   func(args ...string) error
	}{
		{
			name:   "Hash and Compare Valid Password",
			params: []string{"mysecretpassword"},
			exec: func(args ...string) error {
				password := args[0]
				hashedPassword, err := HashPassword(password)
				if err != nil {
					return fmt.Errorf("HashPassword returned error: %v", err)
				}

				err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
				if err != nil {
					return fmt.Errorf("Failed to compare password: %v", err)
				}

				return nil
			},
		},
		{
			name:   "Compare Incorrect Password",
			params: []string{"mysecretpassword", "wrongpassword"},
			exec: func(args ...string) error {
				password := args[0]
				incorrectPassword := args[1]
				hashedPassword, err := HashPassword(password)
				if err != nil {
					return fmt.Errorf("HashPassword() returned error: %v", err)
				}

				if ComparePassword(hashedPassword, incorrectPassword) {
					return fmt.Errorf("expected false, got true for incorrect password")
				}

				return nil
			},
		},
		{
			name:   "Compare Empty Password",
			params: []string{"mysecretpassword", ""},
			exec: func(args ...string) error {
				password := args[0]
				incorrectPassword := args[1]
				hashedPassword, err := HashPassword(password)
				if err != nil {
					return fmt.Errorf("HashPassword() returned error: %v", err)
				}

				if ComparePassword(hashedPassword, incorrectPassword) {
					return fmt.Errorf("expected false, got true for empty password")
				}

				return nil
			},
		},
		{
			name:   "Compare Empty Stored Password",
			params: []string{"mysecretpassword", ""},
			exec: func(args ...string) error {
				password := args[0]
				incorrectPassword := args[1]
				hashedPassword, err := HashPassword(password)
				if err != nil {
					return fmt.Errorf("HashPassword() returned error: %v", err)
				}

				if ComparePassword(hashedPassword, incorrectPassword) {
					return fmt.Errorf("expected false, got true for empty stored password")
				}

				return nil
			},
		},
		{
			name:   "Hash Empty Password",
			params: []string{""},
			exec: func(args ...string) error {
				password := args[0]

				_, err := HashPassword(password)
				if err == nil {
					return fmt.Errorf("expected error, got nil for empty password")
				}
				return nil
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.exec(tt.params...)
			if err != nil {
				t.Errorf("test '%s' failed: %v", tt.name, err)
			}
		})
	}
}
