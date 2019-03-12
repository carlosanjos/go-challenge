package main

import (
	gochallenge "go-challenge/src/go-challenge"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func hello(c echo.Context) error {
	return c.HTML(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	gochallenge.Routes(e)
	e.Logger.Debug(e.Start(os.Getenv("PORT")))
}
