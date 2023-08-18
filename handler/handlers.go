package handler

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Print("IndexHandler")
	t := template.Must(
		template.ParseFiles("web/template/index.gohtml", "web/template/base.gohtml"),
	)

	fmt.Println(t)
	err := t.ExecuteTemplate(w, "base", nil)
	if err != nil {
		panic(err)
	}
}
