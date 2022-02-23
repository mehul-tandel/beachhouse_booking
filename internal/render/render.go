package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/mehul-tandel/beachhouse_booking/internal/config"
	"github.com/mehul-tandel/beachhouse_booking/internal/models"
)

var app *config.AppConfig

var functions = template.FuncMap{}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(p *models.TemplateData, r *http.Request) *models.TemplateData {
	p.CSRFToken = nosurf.Token(r)
	return p
}

// RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {

	var tc map[string]*template.Template
	// Create template cache everytime if in development or else get templates from cache
	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()

	}
	// get the required template
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("String path from handler does not match any template in cache")
	}

	// create a buffer to write the template in it
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	// write the template to the buffer
	_ = t.Execute(buf, td) // plus the template data passed by handler

	// write the template from the buffer to http.ResponseWriter
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}
