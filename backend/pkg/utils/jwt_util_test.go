package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"web-standalone-template/pkg/configs"
)

func TestVerify(t *testing.T) {
	configs.JWTSecretKey = "secret"
	claims := jwt.MapClaims{}
	claims["userID"] = 12345

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(configs.JWTSecretKey))

	assert.NoError(t, err)

	id, err := Verify(signedToken)
	assert.NoError(t, err)
	assert.Equal(t, 12345, id)
}

func TestGenerateToken(t *testing.T) {
	configs.JWTSecretKey = "secret"
	signedToken, err := GenerateToken(12345)
	assert.NoError(t, err)

	token, err := jwt.Parse(signedToken, keyFunc)
	assert.NoError(t, err)

	assert.Equal(t, float64(12345), token.Claims.(jwt.MapClaims)["userID"].(float64))
}
