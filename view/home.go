package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// HomeHandler renders the index view
func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "index.htm", map[string]interface{}{})
}
