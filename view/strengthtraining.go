package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// StrengthtrainingHandler renders the weight view
func StrengthtrainingHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "strengthtraining.htm", map[string]interface{}{})
}
