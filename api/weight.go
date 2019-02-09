package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

// Measurement contains csv/json data
type Measurement struct {
	Date           string `json:"date"`
	Height         string `json:"height"`
	Weight         string `json:"weight"`
	Bmi            string `json:"bmi"`
	BmiOverweight  string `json:"bmioverweight"`
	BmiUnderweight string `json:"bmiunderweight"`
}

// GetWeight returns all weight data
func GetWeight() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		lines := helper.ReadAllDataFromCSV(c.Get("config").(*helper.Config).WeightData)

		var measurements []Measurement
		firstLine := true
		for _, line := range lines {
			data := Measurement{
				Date:           line[0],
				Height:         line[1],
				Weight:         line[2],
				Bmi:            line[3],
				BmiOverweight:  line[4],
				BmiUnderweight: line[5],
			}
			if !firstLine {
				measurements = append(measurements, data)
			}
			firstLine = false
		}

		jsonData := helper.DataToJSON(&measurements)
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// PostWeight saves new weight data
func PostWeight() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		measurement := Measurement{
			Date:           time.Now().Format("2006-01-02 15:04:05"),
			Height:         c.FormValue("height"),
			Weight:         c.FormValue("weight"),
			Bmi:            "",
			BmiOverweight:  "25",
			BmiUnderweight: "18.5",
		}

		// We rely on the client here...
		height, err := strconv.ParseFloat(measurement.Height, 32)
		weight, err := strconv.ParseFloat(measurement.Weight, 32)
		bmi := ((weight * 1.0) / (((height * 0.01) * height) * 0.01))
		measurement.Bmi = strconv.FormatFloat(bmi, 'f', -1, 64)
		fmt.Println(measurement)

		data := []string{measurement.Date, measurement.Height, measurement.Weight, measurement.Bmi, measurement.BmiOverweight, measurement.BmiUnderweight}
		helper.SaveDataToCSV(c.Get("config").(*helper.Config).WeightData, data)
		jsonData := helper.DataToJSON(&[]Measurement{measurement})
		return c.JSONBlob(http.StatusOK, jsonData)
	}
}

// DeleteWeight deletes weight data
func DeleteWeight() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		date := c.Param("date")
		fmt.Println("Delete by date", date)
		if date == "" {
			return c.JSONBlob(http.StatusBadRequest, []byte(`[]`))
		}

		deleted := weightDeleteLine(c.Get("config").(*helper.Config).WeightData, date)
		if !deleted {
			return c.JSONBlob(http.StatusNotFound, []byte(`[]`))
		}
		return c.JSONBlob(http.StatusOK, []byte(`[]`))
	}
}

func weightDeleteLine(filename string, startText string) bool {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(f), "\n")

	newLines := lines[:]
	found := false
	for i, line := range lines {
		if strings.HasPrefix(line, startText) {
			newLines = append(newLines[:i], newLines[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return false
	}

	output := strings.Join(newLines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		panic(err)
	}

	return true
}
