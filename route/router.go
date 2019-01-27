package route

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spech66/lifelogspd/api"
	"github.com/spech66/lifelogspd/view"
)

// Init initializes echo and routes
func Init() *echo.Echo {
	// Echo instance
	e := echo.New()
	//e.Debug()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// View routes => handler
	e.GET("/", view.HomeHandler)

	// Api routes => handler
	apiGroup := e.Group("/api")
	{
		apiGroup.GET("/weight", api.GetWeight())
		//apiGroup.GET("/weight/:date", api.GetWeight())
		//apiGroup.POST("/weight", api.PostWeight())
	}

	return e
}
