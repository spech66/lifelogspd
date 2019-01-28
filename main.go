package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo"
	"github.com/spech66/lifelogspd/route"
)

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

var flagBind string
var flagConfig string

func init() {
	flag.StringVar(&flagBind, "bind", ":8066", "ip:port to bind the server to")
	flag.StringVar(&flagConfig, "config", "config.json", "the config file to read from")
	flag.Parse()
}

func main() {
	fmt.Println("LifelogSPDaemon")
	fmt.Println("Binding server to", flagBind)
	fmt.Println("Reding config from", flagConfig)

	router := route.Init()

	// Parse all templates
	templates := make(map[string]*template.Template)
	templates["index.htm"] = template.Must(template.ParseFiles("templates/index.htm", "templates/_base.htm"))
	templates["weight.htm"] = template.Must(template.ParseFiles("templates/weight.htm", "templates/_base.htm"))
	templates["journal.htm"] = template.Must(template.ParseFiles("templates/journal.htm", "templates/_base.htm"))
	templates["strengthtraining.htm"] = template.Must(template.ParseFiles("templates/strengthtraining.htm", "templates/_base.htm"))
	templates["enduranceworkout.htm"] = template.Must(template.ParseFiles("templates/enduranceworkout.htm", "templates/_base.htm"))
	router.Renderer = &TemplateRegistry{
		templates: templates,
	}

	// Start the server
	router.Logger.Fatal(router.Start(flagBind))
}
