package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// Strengthtraining contains csv/json data
type Strengthtraining struct {
	Date     string  `json:"date"`
	Exercise string  `json:"exercise"`
	Reps     int     `json:"reps"`
	Weight   float64 `json:"weight"`
	Notes    string  `json:"notes"`
	Rating   int     `json:"rating"`
}

// GetStrengthtraining returns all strength training data
func GetStrengthtraining() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}

// GetStrengthtrainingByDate returns strength training data by date
func GetStrengthtrainingByDate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}

// PostStrengthtraining saves new strength training data
func PostStrengthtraining() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}

// DeleteStrengthtraining deletes strength training data
func DeleteStrengthtraining() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}

// GetStrengthtrainingExercises returns all exercises to select for the training
func GetStrengthtrainingExercises() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		exercises := c.Get("config").(*helper.Config).StrengthtrainingExercises
		jsonData := helper.DataToJSON(&exercises)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}
