package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	password := "password"
	hash, err := GeneratePassword(password)
	assert.NoError(t, err)
	assert.True(t, ComparePasswords(hash, password))
	assert.False(t, ComparePasswords(hash, "p4ssw0rd"))
}
