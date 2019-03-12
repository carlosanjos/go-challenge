package gochallenge

import (
	"go-challenge/src/go-challenge/bar"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo) {
	e.GET("/foo", bar.Show)
}
