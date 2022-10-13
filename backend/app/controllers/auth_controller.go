package controllers

import (
	"backend/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Authorize(c echo.Context) (string, error) {
	token, err := utils.RequestLookUp(c, "token", "query", "form")
	if err != nil {
		return "", err
	}
	return utils.Verify(token)
}

func Verify(c echo.Context) error {
	userID, err := Authorize(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.FailResponse(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(echo.Map{
		"user_id": userID,
	}))
}
