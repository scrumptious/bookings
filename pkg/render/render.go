package render

import (
	"bytes"
	"fmt"
	"github.com/scrumptious/bookings/pkg/config"
	"github.com/scrumptious/bookings/pkg/models"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var pageSearchPattern = "./templates/*.page.tmpl.html"
var layoutSearchPattern = "./templates/*.layout.tmpl"
var tc map[string]*template.Template
var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}
func RenderTemplate(rw http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var err error
	if app.UseCache {
		//get the template cache from app config
		tc = app.TemplateCache
	} else {
		//create new template cache
		tc, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("Failed to create template cache")
		}
	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Failed to find cached template ")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	err = t.Execute(buf, td)
	if err != nil {
		fmt.Println(err)
	}

	_, err = buf.WriteTo(rw)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all pages
	pages, err := filepath.Glob(pageSearchPattern)
	if err != nil {
		return nil, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		layouts, err := filepath.Glob(layoutSearchPattern)
		if err != nil {
			return nil, err
		}

		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(layoutSearchPattern)
			if err != nil {
				return nil, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
