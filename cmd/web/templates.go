package main

import (
	"html/template"
	"path/filepath"
)

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	patterns := []string{
		"./assets/html/pages/*.html",
		"./assets/html/pages/errors/*.html",
	}

	var pages []string
	for _, pattern := range patterns {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}
		pages = append(pages, matches...)
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles("./assets/html/base.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob("./assets/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
