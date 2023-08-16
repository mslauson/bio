package handler

import "github.com/gorilla/mux"

func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler)

	return r
}
