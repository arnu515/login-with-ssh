package main

import (
	"flag"
	"strings"

	"login-with-ssh/db"
	"login-with-ssh/templates"
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
		component := pages.Signup(c, nil, nil)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.POST("/signup", func(c echo.Context) error {
		component := pages.Signup(c, nil, nil)

		// convert form data to struct
		form, err := c.FormParams()
		if err != nil {
			templates.FlashMessage(templates.Message{
				Type:    "error",
				Message: "Could not parse form body: " + err.Error(),
			})
			return component.Render(c.Request().Context(), c.Response())
		}
		username := form.Get("username")
		email := form.Get("email")
		sshKey := form.Get("sshKey")

		if strings.Trim(username, " ") == "" ||
			strings.Trim(email, " ") == "" ||
			strings.Trim(sshKey, " ") == "" {
			templates.FlashMessage(templates.Message{
				Type:    "error",
				Message: "Please fill all fields",
			})
			return component.Render(c.Request().Context(), c.Response())
		}

		// validate ssh key
		// key format: ssh-rsa/ssh-ed25519 KEY COMMENT
		keySplit := strings.SplitN(sshKey, " ", 3)

		if keySplit[0] != "ssh-rsa" && keySplit[0] != "ssh-ed25519" {
			templates.FlashMessage(templates.Message{
				Type:    "error",
				Message: "Only RSA and ED25519 keys are allowed.",
			})
			return component.Render(c.Request().Context(), c.Response())
		}

		userId, err := db.SignupUser(c.Request().Context(), db.DB, username, email, keySplit[0], keySplit[1])
		if err != nil {
			templates.FlashMessage(templates.Message{
				Type:    "error",
				Message: "SignupUser: " + err.Error(),
			})
			return component.Render(c.Request().Context(), c.Response())
		}

		authId, err := db.AddToSSHQueue(c.Request().Context(), db.DB, userId)
		if err != nil {
			templates.FlashMessage(templates.Message{
				Type:    "error",
				Message: "AddToSSHQueue: " + err.Error(),
			})
			return component.Render(c.Request().Context(), c.Response())
		}

		// temp
		component = pages.Signup(c, &authId, &c.Request().Host)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.GET("/ssh-key", func(c echo.Context) error {
		component := pages.SshKeyPage(c)
		return component.Render(c.Request().Context(), c.Response())
	})

	app.GET("/auth/:id", func(c echo.Context) error {
		_ = c.Param("id")
		// temp
		if false {
			return c.NoContent(204)
		}
		return c.String(401, "Unauthorized")
	})

	app.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]bool{"ok": true})
	})

	app.Logger.Fatal(app.Start(":8080"))
}
