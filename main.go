package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}

type PageData struct {
	Title string
	Todos []Todo
}

func todo(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Kevin Arellano",
		Todos: []Todo{
			{Item: "Install Go", Done: true},
			{Item: "Learn Go", Done: false},
			{Item: "Yup", Done: false},
		},
	}
	tmpl.Execute(w, data)
}

func about(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "About",
		Todos: []Todo{
			{Item: "Welcome to the about page", Done: true},
		},
	}
	tmpl.Execute(w, data)
}

func main() {
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	mux.HandleFunc("/about", about)
	mux.HandleFunc("/", todo)

	log.Fatal(http.ListenAndServe(":3000", mux))
}
