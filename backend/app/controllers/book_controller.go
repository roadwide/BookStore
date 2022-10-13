package controllers

import (
	"backend/app/models"
	"backend/app/queries"
	"backend/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddBook(c echo.Context) error {
	userID, err := Authorize(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, utils.FailResponse(err.Error(), nil))

	}
	req := &models.AddBookRequest{}
	err = c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}
	if req.UserID != userID {
		return c.JSON(http.StatusUnauthorized, utils.FailResponse("Illegal Access", nil))
	}
	book, err := queries.DataBase.CrateBook(req.UserID, req.Name, req.PicURL, req.Price)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(book))
}

func GetBook(c echo.Context) error {
	req := &models.GetBookRequest{}
	err := c.Bind(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	if req.UserID == "" {
		books, err := queries.DataBase.GetAllBook()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
		}
		return c.JSON(http.StatusOK, utils.SuccessResponse(books))
	}
	books, err := queries.DataBase.GetBook(req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse(books))
}
