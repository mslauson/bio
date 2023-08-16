package handler

import (
	"log"
	"net/http"

	"gitea.slauson.io/blog/blog-ui/components/base"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Default().Print("IndexHandler")
	err := base.Base().Render(r.Context(), w)
	if err != nil {
		panic(err)
	}
}
