package handler

import (
	"net/http"

	"gitea.slauson.io/blog/blog-ui/components/base"
	"github.com/a-h/templ"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(base.Base())
}
