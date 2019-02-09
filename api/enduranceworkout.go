package api

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// Enduranceworkout contains csv/json data
type Enduranceworkout struct {
	Date     string `json:"date"`
	Exercise string `json:"exercise"`
	Distance int    `json:"distance"`
	Duration int    `json:"duration"`
	Notes    string `json:"notes"`
	Rating   int    `json:"rating"`
}

// GetEnduranceworkout returns all endurance workout data
func GetEnduranceworkout() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
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
