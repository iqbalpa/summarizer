package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt the password")
	}
	return string(encrypted), err
}

func ComparePassword(password, encryptedPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPass), []byte(password))
	if err != nil {
		return false, fmt.Errorf("password incorrect")
	}
	return true, nil
}
