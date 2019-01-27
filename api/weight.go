package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

// GetWeight returns all weight data
func GetWeight() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(
			http.StatusOK,
			[]byte(`[{ "date": "20190101", "weight": "88.0", "height": "192" }, { "date": "20190110", "weight": "89.0", "height": "192" }]`),
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
