package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

// Create a struct that holds information to be displayed in our HTML file
type Welcome struct {
	Name string
	Time string
}

// Go application entrypoint
func main() {
	// Instantiate a Welcome struct object and pass in some random information.
	// We shall get the name of the user as a query parameter from the URL
	// welcome := Welcome{"Anonymous", time.Now().Format(time.Stamp)}
	//
	// // We tell Go exactly where we can find our html file. We ask Go to parse the html file (Notice
	// // the relative path). We wrap it in a call to template.Must() which handles any errors and halts if there are fatal errors
	//
	// templates := template.Must(template.ParseGlob("web/template/*.go.tmpl"))
	// templates = template.Must(templates.ParseGlob("web/template/pages/*.go.tmpl"))
	//
	// // Our HTML comes with CSS that go needs to provide when we run the app. Here we tell go to create
	// // a handle that looks in the static directory, go then uses the "/static/" as a url that our
	// // html can refer to when looking for our css and other files.
	//
	// http.Handle(
	// 	"/static/",
	// 	http.StripPrefix("/static/", http.FileServer(http.Dir("static"))),
	// ) // Go looks in the relative "static" directory first using http.FileServer(), then matches it to a

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	if name := r.FormValue("name"); name != "" {
	// 		welcome.Name = name
	// 	}
	//
	// 	if err := templates.ExecuteTemplate(w, "pages/home", welcome); err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	}
	// })

	http.Handle("/", templ.Handler(Index()))

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
