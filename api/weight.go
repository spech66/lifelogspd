package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/helper"
)

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
		weightPath := c.Get("config").(*helper.Config).WeightPath
		year := time.Now().Format("2006")
		filename := filepath.Join(weightPath, year+"_weight.csv")
		fmt.Println("Weight data from", filename)

		f, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		r := csv.NewReader(f)
		r.Comma = ';'
		lines, err := r.ReadAll()
		if err != nil {
			panic(err)
		}

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

		var jsonData []byte
		jsonData, err = json.Marshal(&measurements)
		if err != nil {
			panic(err)
		}

		return c.JSONBlob(
			http.StatusOK,
			jsonData,
		)
	}
}

// PostWeight saves new weight data
func PostWeight() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(
			http.StatusOK,
			[]byte(`{ "date": "20190201", "weight": "90.0", "height": "192" }`),
		)
	}
}

// DeleteWeight deletes weight data
func DeleteWeight() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		date := c.QueryParam("date")
		fmt.Println("Date ", date)
		return c.JSONBlob(
			http.StatusOK,
			[]byte(`[{ "date": "20190101", "weight": "88.0", "height": "192" }, { "date": "20190110", "weight": "89.0", "height": "192" }]`),
		)
	}
}
