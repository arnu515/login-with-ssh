package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"login-with-ssh/templates"
)

func main() {
	app := echo.New()

	app.Static("/", "./static")

	app.GET("/", func(c echo.Context) error {
		component := templates.Index()
		return component.Render(context.Background(), c.Response())
	})

	app.Logger.Fatal(app.Start(":8080"))
}
