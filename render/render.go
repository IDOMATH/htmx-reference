package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/idomath/htmx-reference/types"
)

var views = "./views"

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.html", views))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.html", views))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.html", views))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// Template renders html templates for pages
func Template(w http.ResponseWriter, r *http.Request, tmpl string, data *types.TemplateData) error {
	var tc map[string]*template.Template

	tc, err := CreateTemplateCache()
	if err != nil {
		return err
	}

	// Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		return fmt.Errorf("can't get template %s from cache", tmpl)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, data)
	if err != nil {
		log.Fatal(err)
	}

	// Render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
		return err
	}

	return nil
}
