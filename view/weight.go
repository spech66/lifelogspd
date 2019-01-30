package view

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// WeightHandler renders the weight view
func WeightHandler(c echo.Context) error {
	defaultHeight := c.Get("config").(*helper.Config).DefaultHeight
	return c.Render(http.StatusOK, "weight.htm", map[string]interface{}{
		"height": defaultHeight,
	})
}
