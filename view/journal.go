package view

import (
	"net/http"

	"github.com/labstack/echo"
)

// JournalHandler renders the weight view
func JournalHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "journal.htm", map[string]interface{}{})
}
