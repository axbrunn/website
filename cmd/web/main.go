package main

import (
	"fmt"
	"os"
)

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

func main() {
	p1 := &Page{
		Title: "Healthcheck",
		Body:  []byte("I'm a healthy Page!"),
	}
	p1.save()
	p2, _ := loadPage("Healthcheck")
	fmt.Println(string(p2.Body))
}
