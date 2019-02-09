package api

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// Enduranceworkout contains csv/json data
type Enduranceworkout struct {
	Date     string `json:"date"`
	Exercise string `json:"exercise"`
	Distance int64  `json:"distance"`
	Duration string `json:"duration"`
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
			distance, _ := strconv.ParseInt(line[2], 10, 64)
			rating, _ := strconv.ParseInt(line[5], 10, 64)

			data := Enduranceworkout{
				Date:     line[0],
				Exercise: line[1],
				Distance: distance,
				Duration: line[3],
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

/*// GetEnduranceworkoutByDate returns endurance workout data by date
func GetEnduranceworkoutByDate() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
	}
}*/

// PostEnduranceworkout saves new endurance workout data
func PostEnduranceworkout() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// beware that we drop the err here as this is only a internal server solution!
		distance, _ := strconv.ParseInt(c.FormValue("distance"), 10, 64)
		rating, _ := strconv.ParseInt(c.FormValue("rating"), 10, 64)

		// Add seconds component if not send by browser
		duration := c.FormValue("duration")
		if len(duration) == 5 {
			duration = duration + ":00"
		}

		enduranceworkout := Enduranceworkout{
			Date:     time.Now().Format("2006-01-02 15:04:05"),
			Exercise: c.FormValue("exercise"),
			Distance: distance,
			Duration: c.FormValue("duration"),
			Notes:    c.FormValue("notes"),
			Rating:   rating,
		}

		fmt.Println(enduranceworkout)

		data := []string{
			enduranceworkout.Date,
			enduranceworkout.Exercise,
			strconv.FormatInt(enduranceworkout.Distance, 10),
			enduranceworkout.Duration,
			enduranceworkout.Notes,
			strconv.FormatInt(enduranceworkout.Rating, 10),
		}
		helper.SaveDataToCSV(c.Get("config").(*helper.Config).EnduranceworkoutData, data)
		jsonData := helper.DataToJSON(&[]Enduranceworkout{enduranceworkout})
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// DeleteEnduranceworkout deletes endurance workout data
func DeleteEnduranceworkout() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		date := c.Param("date")
		fmt.Println("Delete by date", date)
		if date == "" {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		deleted := helper.DeleteLineFromSCV(c.Get("config").(*helper.Config).EnduranceworkoutData, date)
		if !deleted {
			return c.JSONBlob(http.StatusNotFound, []byte(`[]`))
		}
		return c.JSONBlob(http.StatusOK, []byte(`[]`))
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
