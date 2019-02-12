package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// StrengthtrainingGraphHandler renders the strengthtraining graph view
func StrengthtrainingGraphHandler(c echo.Context) error {
	exercise := c.Param("exercise")
	return c.Render(http.StatusOK, "strengthtraininggraph.htm", map[string]interface{}{
		"exercise": exercise,
	})
}
