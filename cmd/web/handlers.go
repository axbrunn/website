package main

import (
	"net/http"
)

func (app *application) errorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	ErrorData := struct {
		Status  int
		Message string
	}{
		Status:  status,
		Message: message,
	}

	err := app.renderTemplate(w, status, "error", ErrorData)
	if err != nil {
		app.logger.Error(err.Error(), "method", r.Method, "uri", r.URL.RequestURI())
		http.Error(w, http.StatusText(status), status)
	}
}

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, http.StatusOK, "home", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) fooHandler(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, http.StatusOK, "foo", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *application) bazHandler(w http.ResponseWriter, r *http.Request) {
	err := app.renderTemplate(w, http.StatusOK, "baz", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}
