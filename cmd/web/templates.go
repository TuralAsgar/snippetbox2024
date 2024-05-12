package main

import (
	"io/fs"
	"path/filepath"
	"snippetbox.turalasgar.com/ui"
	"text/template"
	"time"

	"snippetbox.turalasgar.com/internal/models"
)

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

type templateData struct {
	CurrentYear     int
	Snippet         models.Snippet
	Snippets        []models.Snippet
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
	User            models.User
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	//pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html")
	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl.html",
			"html/partials/*.tmpl.html",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		//ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.tmpl.html")
		//if err != nil {
		//	return nil, err
		//}

		//ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		//if err != nil {
		//	return nil, err
		//}
		//
		//ts, err = ts.ParseFiles(page)
		//if err != nil {
		//	return nil, err
		//}

		cache[name] = ts
	}

	return cache, nil
}
