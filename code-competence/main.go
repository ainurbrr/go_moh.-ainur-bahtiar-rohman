package main

import (
	"code-competence/config"
	"code-competence/middlewares"
	"code-competence/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.Init()
	e := echo.New()

	routes.Routes(e, db)
	middlewares.LogMiddleware(e)

	e.Logger.Fatal(e.Start(":8080"))
}
