package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// Strengthtraining contains csv/json data
type Strengthtraining struct {
	Date     string  `json:"date"`
	Exercise string  `json:"exercise"`
	Reps     int64   `json:"reps"`
	Weight   float64 `json:"weight"`
	Notes    string  `json:"notes"`
	Rating   int64   `json:"rating"`
}

// GetStrengthtraining returns all strength training data
func GetStrengthtraining() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		lines := helper.ReadAllDataFromCSV(c.Get("config").(*helper.Config).StrengthtrainingData)

		var strengthtrainings []Strengthtraining
		firstLine := true
		for _, line := range lines {
			reps, _ := strconv.ParseInt(line[2], 0, 64)
			weight, _ := strconv.ParseFloat(line[3], 64)
			rating, _ := strconv.ParseInt(line[5], 0, 64)

			data := Strengthtraining{
				Date:     line[0],
				Exercise: line[1],
				Reps:     reps,
				Weight:   weight,
				Notes:    line[4],
				Rating:   rating,
			}
			if !firstLine {
				strengthtrainings = append(strengthtrainings, data)
			}
			firstLine = false
		}

		jsonData := helper.DataToJSON(&strengthtrainings)
		return c.JSONBlob(http.StatusOK, jsonData)
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
