package main

import (
	"fmt"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./assets/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.errorHandler(w, r, http.StatusNotFound, fmt.Sprintf("The page %s were you looking for does not exist.", r.URL.RequestURI()))
	})
	mux.HandleFunc("GET /{$}", app.homeHandler)
	mux.HandleFunc("GET /foo", app.fooHandler)
	mux.HandleFunc("GET /baz", app.bazHandler)

	return mux
}
