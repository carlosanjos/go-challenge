package bar

import (
	"net/http"

	"github.com/labstack/echo"
)

func Show(c echo.Context) error {
	return c.HTML(http.StatusOK, "<strong>bar</strong>")
}
