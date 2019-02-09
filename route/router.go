package route

import (
	"errors"
	"io"
	"text/template"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spech66/lifelogspd/api"
	"github.com/spech66/lifelogspd/helper"
	"github.com/spech66/lifelogspd/view"
)

func configExtender(config *helper.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("config", config)
			return next(c)
		}
	}
}

// TemplateRegistry defines the template registry struct
type TemplateRegistry struct {
	templates map[string]*template.Template
}

// Render templates by implementing e.Renderer interface => https://blog.boatswain.io/post/setup-nested-html-template-in-go-echo-web-framework/
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := t.templates[name]

	if !ok {
		err := errors.New("Template not found -> " + name)
		return err
	}

	return tmpl.ExecuteTemplate(w, "_base.htm", data)
}

// Init initializes echo and routes => https://github.com/eurie-inc/echo-sample
func Init(config *helper.Config) *echo.Echo {
	// Echo instance
	e := echo.New()
	//e.Debug()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(configExtender(config))

	// View routes => handler
	e.GET("/", view.HomeHandler)
	e.GET("/weight", view.WeightHandler)
	e.GET("/weightgraph", view.WeightGraphHandler)
	e.GET("/journal", view.JournalHandler)
	e.GET("/strengthtraining", view.StrengthtrainingHandler)
	e.GET("/enduranceworkout", view.EnduranceworkoutHandler)

	// Parse all templates
	templates := make(map[string]*template.Template)
	templates["index.htm"] = template.Must(template.ParseFiles("templates/index.htm", "templates/_base.htm"))
	templates["weight.htm"] = template.Must(template.ParseFiles("templates/weight.htm", "templates/_base.htm"))
	templates["weightgraph.htm"] = template.Must(template.ParseFiles("templates/weightgraph.htm", "templates/_base.htm"))
	templates["journal.htm"] = template.Must(template.ParseFiles("templates/journal.htm", "templates/_base.htm"))
	templates["strengthtraining.htm"] = template.Must(template.ParseFiles("templates/strengthtraining.htm", "templates/_base.htm"))
	templates["enduranceworkout.htm"] = template.Must(template.ParseFiles("templates/enduranceworkout.htm", "templates/_base.htm"))
	e.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Api routes => handler
	apiGroup := e.Group("/api")
	{
		apiGroup.GET("/weight", api.GetWeight())
		//apiGroup.GET("/weight/:date", api.GetWeightByDate())
		apiGroup.POST("/weight", api.PostWeight())
		apiGroup.DELETE("/weight/:date", api.DeleteWeight())

		apiGroup.GET("/journal", api.GetJournal())
		apiGroup.GET("/journal/:date", api.GetJournalByDate())
		apiGroup.POST("/journal", api.PostJournal())
		apiGroup.DELETE("/journal/:date", api.DeleteJournal())

		apiGroup.GET("/strengthtraining", api.GetStrengthtraining())
		//apiGroup.GET("/strengthtraining/:date", api.GetStrengthtrainingByDate())
		apiGroup.POST("/strengthtraining", api.PostStrengthtraining())
		apiGroup.DELETE("/strengthtraining/:date", api.DeleteStrengthtraining())
		apiGroup.GET("/strengthtraining/exercises", api.GetStrengthtrainingExercises())

		apiGroup.GET("/enduranceworkout", api.GetEnduranceworkout())
		//apiGroup.GET("/enduranceworkout/:date", api.GetEnduranceworkoutByDate())
		apiGroup.POST("/enduranceworkout", api.PostEnduranceworkout())
		apiGroup.DELETE("/enduranceworkout/:date", api.DeleteEnduranceworkout())
		apiGroup.GET("/enduranceworkout/exercises", api.GetEnduranceworkoutExercises())
	}

	return e
}
