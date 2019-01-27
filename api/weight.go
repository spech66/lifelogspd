package api

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetWeight returns all weight data
func GetWeight() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSONBlob(
			http.StatusOK,
			[]byte(`{ "date": "20190101", "weight": "88.0", "height": "192" }`),
		)
	}
}
