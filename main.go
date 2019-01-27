package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/route"
)

// TemplateRegistry definition
type TemplateRegistry struct {
	templates *template.Template
}

// Render templates by implementing e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	fmt.Println("LifelogSPDaemon")

	router := route.Init()

	// Parse all templates
	router.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("templates/*.htm")),
	}

	router.Logger.Fatal(router.Start(":8066"))
}
