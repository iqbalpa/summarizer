package utils

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var (
	key []byte
)

func init() {
	key = []byte(os.Getenv("JWT_KEY"))
}

func GenerateToken(username, id string) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"iss":      "iqbalpa-summarizer",
		"username": username,
		"userId":   id,
	})
	s, err := t.SignedString(key)
	if err != nil {
		return "", err
	}
	return s, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %v", err)
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token claims")
}
