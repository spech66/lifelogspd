package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// Enduranceworkout contains csv/json data
type Enduranceworkout struct {
	Date     string `json:"date"`
	Exercise string `json:"exercise"`
	Distance int64  `json:"distance"`
	Duration int64  `json:"duration"`
	Notes    string `json:"notes"`
	Rating   int64  `json:"rating"`
}

// GetEnduranceworkout returns all endurance workout data
func GetEnduranceworkout() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		lines := helper.ReadAllDataFromCSV(c.Get("config").(*helper.Config).EnduranceworkoutData)

		var enduranceworkouts []Enduranceworkout
		firstLine := true
		for _, line := range lines {
			distance, _ := strconv.ParseInt(line[2], 0, 64)
			duration, _ := strconv.ParseInt(line[3], 0, 64)
			rating, _ := strconv.ParseInt(line[5], 0, 64)

			data := Enduranceworkout{
				Date:     line[0],
				Exercise: line[1],
				Distance: distance,
				Duration: duration,
				Notes:    line[4],
				Rating:   rating,
			}
			if !firstLine {
				enduranceworkouts = append(enduranceworkouts, data)
			}
			firstLine = false
		}

		jsonData := helper.DataToJSON(&enduranceworkouts)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// GetEnduranceworkoutByDate returns endurance workout data by date
func GetEnduranceworkoutByDate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}

// PostEnduranceworkout saves new endurance workout data
func PostEnduranceworkout() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}

// DeleteEnduranceworkout deletes endurance workout data
func DeleteEnduranceworkout() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}

// GetEnduranceworkoutExercises returns all exercises to select for the training
func GetEnduranceworkoutExercises() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		exercises := c.Get("config").(*helper.Config).EnduranceworkoutExercises
		jsonData := helper.DataToJSON(&exercises)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}
