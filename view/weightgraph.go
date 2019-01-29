package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// WeightGraphHandler renders the weight graph view
func WeightGraphHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "weightgraph.htm", map[string]interface{}{})
}
