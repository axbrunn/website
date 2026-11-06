package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./assets/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.notFound(w)
	})
	mux.HandleFunc("GET /{$}", app.homeHandler)
	mux.HandleFunc("GET /foo", app.fooHandler)

	return mux
}
