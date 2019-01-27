package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// WeightHandler renders the weight view
func WeightHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "weight.htm", map[string]interface{}{
		"height": "192.0",
	})
}
