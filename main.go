package main

import (
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

// Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

var templates *template.Template

func main() {
	// Parse your templates
	templates = template.Must(template.ParseGlob("web/template/*.gohtml"))

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type HomeData struct {
	Name string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := HomeData{
		Name: "Matthew Slauson",
	}

	err := templates.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
