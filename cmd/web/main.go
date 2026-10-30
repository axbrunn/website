package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

const ADDR = ":8080"

type Page struct {
	Title string
	Body  []byte
}

func constructFileName(name, ext string) string {
	return name + ext
}

func (p *Page) save() error {
	filename := constructFileName(p.Title, ".txt")
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := constructFileName(title, ".txt")
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/html/index.html")
	t.Execute(w, nil)
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		t, _ := template.ParseFiles("assets/html/notFound.html")
		t.Execute(w, nil)
	}
}

func main() {
	p1 := &Page{
		Title: "test",
		Body:  []byte("I'm a healthy Page!"),
	}
	p1.save()

	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHandler)

	log.Fatal(http.ListenAndServe(ADDR, nil))
}
