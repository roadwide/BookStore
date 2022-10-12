package routes

import "github.com/labstack/echo/v4"

func StaticRoutes(a *echo.Echo) {
	a.Static("/", "www")
	a.Static("/img", "img")
}
