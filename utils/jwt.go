package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var SecretKey = "SECRET_TOKEN"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	return webtoken, nil
}

func VerifyToken(tokenJWT string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenJWT, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Unexpected: %v", t.Header["alg"])
		}

		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func DecodeToken(token string) (jwt.MapClaims, error) {
	t, err := VerifyToken(token)

	if err != nil {
		return nil, err
	}

	claims, isOk := t.Claims.(jwt.MapClaims)

	if isOk && t.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}