# Login with SSH

Proof-of-Concept for logging in with SSH key.

This is a golang echo application which demonstrates sign-in with an SSH key.

In the future, I may add other methods, such as login with magic link, TOTP, passkeys, etc. This repository may be renamed once those are added.

For more information, check out: [LINK TO BE ADDED]

## Tech stack

- [Labstack Echo](https://echo.labstack.com)
- [Templ](https://templ.guide)
- [PicoCSS](https://picocss.com)

## How to run

```bash
$ templ generate

$ go run main.go
```

## Development

Please format your code (`go fmt` and `templ fmt .`) before sending a PR!

Run `templ generate -watch` to automatically generate templates when you save a file.

I'll probably add something like [air](https://github.com/cosmtrek/air) in the future.
