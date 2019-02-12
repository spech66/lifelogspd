package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// EnduranceworkoutHandler renders the weight view
func EnduranceworkoutHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "enduranceworkout.htm", map[string]interface{}{})
}
