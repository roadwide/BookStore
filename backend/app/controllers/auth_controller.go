package controllers

import (
	"backend/pkg/utils"
	"github.com/labstack/echo/v4"
)

func Authorize(c echo.Context) (string, error) {
	token, err := utils.RequestLookUp(c, "token", "query", "form")
	if err != nil {
		return "", err
	}
	return utils.Verify(token)
}
