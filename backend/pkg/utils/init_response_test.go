package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessResponse(t *testing.T) {
	assert.Equal(t, Response{
		Code: 0,
		Msg:  "OK",
		Resp: []string{"hello", "world"},
	}, SuccessResponse([]string{"hello", "world"}))
}

func TestFailResponse(t *testing.T) {
	assert.Equal(t, Response{
		Code: 1,
		Msg:  "error",
		Resp: "not ok",
	}, FailResponse("error", "not ok"))
	assert.Equal(t, Response{
		Code: 1,
		Msg:  "error",
		Resp: nil,
	}, FailResponse("error", nil))
}

func TestDIYResponse(t *testing.T) {
	type node struct {
		a string
		b int
	}
	assert.Equal(t, Response{
		Code: 9,
		Msg:  "test",
		Resp: node{"nah", 9},
	}, DIYResponse(9, "test", node{"nah", 9}))
	assert.NotEqual(t, Response{
		Code: 9,
		Msg:  "test",
		Resp: node{"nah", 9},
	}, DIYResponse(9, "test", node{"nah", 99}))
}
