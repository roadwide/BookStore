package main

import (
	"backend/pkg/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()

	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	app.Use(middleware.CORS())

	routes.PublicRoutes(app)
	routes.GeneralRoutes(app)
	routes.StaticRoutes(app)

	app.Logger.Fatal(app.Start(":8081"))
}
