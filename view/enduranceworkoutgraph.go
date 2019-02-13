package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// EnduranceworkoutGraphHandler renders the enduranceworkout graph view
func EnduranceworkoutGraphHandler(c echo.Context) error {
	exercise := c.Param("exercise")
	return c.Render(http.StatusOK, "enduranceworkoutgraph.htm", map[string]interface{}{
		"exercise": exercise,
	})
}
