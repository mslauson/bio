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
	templates = template.Must(template.ParseGlob("web/template/*.tmpl"))

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{} // Fill in data as required
	err := templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
