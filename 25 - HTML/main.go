package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type user struct {
	Name  string
	Email string
}

func main() {
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		u := user{"Livia", "liviadevolder@gmail.com"}
		
		templates.ExecuteTemplate(w, "home.html", u)
	})

	log.Fatal(http.ListenAndServe(":5005", nil))
}
