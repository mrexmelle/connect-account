package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-iam/authx/session"
)

func main() {
	r := chi.NewRouter()

	r.Route("/sessions", func(r chi.Router) {
		r.Post("/", session.Post)
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
