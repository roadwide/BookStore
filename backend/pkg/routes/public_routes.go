package routes

import (
	"backend/app/controllers"
	"github.com/labstack/echo/v4"
)

func PublicRoutes(a *echo.Echo) {
	userGroup := a.Group("/user")
	userGroup.POST("/register", controllers.Register)
	userGroup.POST("/login", controllers.Login)

	bookGroup := a.Group("/book")
	bookGroup.POST("/add", controllers.AddBook)
	bookGroup.GET("/info", controllers.GetBook)
}
