package controllers

import (
	"backend/app/models"
	"backend/app/queries"
	"backend/pkg/repository"
	"backend/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	req := &models.UserRequest{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error(), nil))
	}

	if req.Name == "" || len(req.Name) > repository.MaxUsernameLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Username", nil))
	}

	if req.Password == "" || len(req.Password) > repository.MaxUserPasswordLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Password", nil))
	}

	hash, err := utils.GeneratePassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error(), nil))
	}

	_, err = queries.DataBase.CrateUser(req.Name, hash, req.Email)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error(), nil))
	}

	token, err := utils.GenerateToken(req.Name)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(&models.UserResp{
		Name:  req.Name,
		Token: token,
	}))
}

func Login(c echo.Context) error {
	req := &models.UserRequest{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error(), nil))
	}

	if req.Name == "" || len(req.Name) > repository.MaxUsernameLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Username", nil))
	}

	if req.Password == "" || len(req.Password) > repository.MaxUserPasswordLength {
		return c.JSON(http.StatusOK, utils.FailResponse("Illegal Password", nil))
	}

	hash, err := queries.DataBase.GetUserHash(req.Name)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error(), nil))
	}

	if !utils.ComparePasswords(hash, req.Password) {
		return c.JSON(http.StatusOK, utils.FailResponse("Username or Password is Incorrect", nil))
	}

	token, err := utils.GenerateToken(req.Name)
	if err != nil {
		return c.JSON(http.StatusOK, utils.FailResponse(err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse(&models.UserResp{
		Name:  req.Name,
		Token: token,
	}))
}
