package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

const ADDR = ":6767"

var validSlug = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*$`)

type application struct {
	templates map[string]*template.Template
}

func (app *application) renderTemplate(w http.ResponseWriter, status int, tmpl string, data any) error {
	ts, ok := app.templates[tmpl+".html"]
	if !ok {
		return fmt.Errorf("the template %s does not exist in the cache", tmpl)
	}

	w.WriteHeader(status)

	return ts.ExecuteTemplate(w, "base", data)
}

func (app *application) notFound(w http.ResponseWriter) {
	err := app.renderTemplate(w, http.StatusNotFound, "notfound", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, http.StatusOK, "home", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) fooHandler(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, http.StatusOK, "foo", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	templateCache, err := newTemplateCache()
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	app := &application{
		templates: templateCache,
	}

	err = run(app)
	if err != nil {
		log.Printf("startup err: %v", err)
		os.Exit(1)
	}
}

func run(app *application) error {

	log.Printf("starting server on %s", ADDR)

	err := http.ListenAndServe(ADDR, app.routes())
	if err != nil {
		return err
	}

	return nil
}
