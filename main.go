package main

import (
	"context"

	"github.com/labstack/echo/v4"
	"login-with-ssh/templates"
	"login-with-ssh/templates/pages"
)

func main() {
	app := echo.New()

	app.Static("/", "./static")

	app.GET("/", func(c echo.Context) error {
		component := pages.Index(c)
		return component.Render(context.Background(), c.Response())
	})

	app.GET("/login", func(c echo.Context) error {
		component := pages.Login(c, nil)
		return component.Render(context.Background(), c.Response())
	})

	app.GET("/signup", func(c echo.Context) error {
		component := pages.Signup(c, nil)
		return component.Render(context.Background(), c.Response())
	})

	app.GET("/ssh-key", func(c echo.Context) error {
		component := pages.SshKeyPage(c)
		return component.Render(context.Background(), c.Response())
	})

	app.Logger.Fatal(app.Start(":8080"))
}
