package utils

import (
	"backend/pkg/configs"
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

func Verify(signedToken string) (userID string, err error) {
	token, err := jwt.Parse(signedToken, keyFunc)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	id, ok := claims["userID"].(string)
	if !ok {
		return "", errors.New("invalid userID")
	}

	return id, nil
}

func GenerateToken(userID string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.JWTSecretKey))
}

func keyFunc(*jwt.Token) (any, error) {
	return []byte(configs.JWTSecretKey), nil
}
