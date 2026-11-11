package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) renderTemplate(w http.ResponseWriter, status int, tmpl string, data any) error {
	ts, ok := app.templates[tmpl+".html"]
	if !ok {
		return fmt.Errorf("the template %s does not exist in the cache", tmpl)
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		return err
	}

	w.WriteHeader(status)

	_, err = buf.WriteTo(w)
	return err
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(err.Error(), "method", method, "uri", uri)
	app.errorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
}
