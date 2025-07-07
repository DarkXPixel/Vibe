package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAuthJWTToken(phone string, secret string, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"sub":  phone,
		"exp":  time.Now().Add(ttl).Unix(),
		"iat":  time.Now().Unix(),
		"type": "verify",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateAuthJWTToken(tokenString, secret string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || claims["type"] != "verify" {
		return "", fmt.Errorf("invalid token type")
	}

	phone, ok := claims["sub"].(string)
	if !ok {
		return "", fmt.Errorf("invalid phone in token")
	}
	return phone, nil
}
