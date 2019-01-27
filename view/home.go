package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// HomeHandler renders the index route
func HomeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "_base.htm", map[string]interface{}{
		"name": "Home",
	})
}
