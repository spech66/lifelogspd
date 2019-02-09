package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// GetStrengthtrainingExercises returns all exercises to select for the training
func GetStrengthtrainingExercises() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		exercises := c.Get("config").(*helper.Config).StrengthtrainingExercises
		jsonData := helper.DataToJSON(&exercises)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}
