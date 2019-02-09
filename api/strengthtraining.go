package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
			reps, _ := strconv.ParseInt(line[2], 10, 64)
			weight, _ := strconv.ParseFloat(line[3], 64)
			rating, _ := strconv.ParseInt(line[5], 10, 64)

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
		// beware that we drop the err here as this is only a internal server solution!
		reps, _ := strconv.ParseInt(c.FormValue("reps"), 10, 64)
		weight, _ := strconv.ParseFloat(c.FormValue("weight"), 64)
		rating, _ := strconv.ParseInt(c.FormValue("rating"), 10, 64)

		strengthtraining := Strengthtraining{
			Date:     time.Now().Format("2006-01-02 15:04:05"),
			Exercise: c.FormValue("exercise"),
			Reps:     reps,
			Weight:   weight,
			Notes:    c.FormValue("notes"),
			Rating:   rating,
		}

		fmt.Println(strengthtraining)

		data := []string{
			strengthtraining.Date,
			strengthtraining.Exercise,
			strconv.FormatInt(strengthtraining.Reps, 10),
			strconv.FormatFloat(strengthtraining.Weight, 'f', -2, 64),
			strengthtraining.Notes,
			strconv.FormatInt(strengthtraining.Rating, 10),
		}
		helper.SaveDataToCSV(c.Get("config").(*helper.Config).StrengthtrainingData, data)
		jsonData := helper.DataToJSON(&[]Strengthtraining{strengthtraining})
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// DeleteStrengthtraining deletes strength training data
func DeleteStrengthtraining() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		date := c.Param("date")
		fmt.Println("Delete by date", date)
		if date == "" {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		deleted := helper.DeleteLineFromSCV(c.Get("config").(*helper.Config).StrengthtrainingData, date)
		if !deleted {
			return c.JSONBlob(http.StatusNotFound, []byte(`[]`))
		}
		return c.JSONBlob(http.StatusOK, []byte(`[]`))
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
