package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// GetEnduranceworkoutExercises returns all exercises to select for the training
func GetEnduranceworkoutExercises() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		exercises := c.Get("config").(*helper.Config).EnduranceworkoutExercises
		jsonData := helper.StringListToJSON(&exercises)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}
