package main

import (
	"flag"

	"login-with-ssh/db"
	"login-with-ssh/templates/pages"

	"github.com/labstack/echo/v4"
)

func init() {
	flag.Parse()
}

func main() {
	defer func() {
		err := db.DB.Close()
		if err != nil {
			panic(err)
		}
	}()

	if ToSeed {
		err := seed()
		if err != nil {
			panic(err)
		}
		return
	}

	app := echo.New()

	app.Static("/", "./static")

	app.GET("/", func(c echo.Context) error {
		component := pages.Index(c)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.GET("/login", func(c echo.Context) error {
		component := pages.Login(c, nil)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.GET("/signup", func(c echo.Context) error {
		component := pages.Signup(c, nil)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.POST("/signup", func(c echo.Context) error {
		component := pages.Signup(c, nil)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.GET("/ssh-key", func(c echo.Context) error {
		component := pages.SshKeyPage(c)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.Logger.Fatal(app.Start(":8080"))
}
