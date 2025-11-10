package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var (
	key []byte
)

func init() {
	key = []byte(os.Getenv("JWT_KEY"))
}

func GenerateToken(username string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":      "iqbalpa-summarizer",
		"username": username,
	})
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}
