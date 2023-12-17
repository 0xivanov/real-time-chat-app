package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashString(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), 14)
	if err != nil {
		return "", fmt.Errorf("failed to hash string: %v", err)
	}
	return string(hashed), nil
}

func CheckHashedString(str string, hashedStr string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(str))
}
