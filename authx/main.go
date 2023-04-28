package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-iam/authx/account"
	"github.com/mrexmelle/connect-iam/authx/config"
	"github.com/mrexmelle/connect-iam/authx/session"
	"go.uber.org/dig"
)

func Config() *config.Config {
	cfg, err := config.New(
		"authx", "yaml",
		[]string{
			"/etc/conf",
			"./config",
		},
	)
	if err != nil {
		panic(err)
	}
	return &cfg
}

func main() {
	container := dig.New()
	container.Provide(Config)
	container.Provide(account.NewService)
	container.Provide(session.NewService)
	container.Provide(account.NewController)
	container.Provide(session.NewController)

	process := func(
		accountController *account.Controller,
		sessionController *session.Controller,
		config *config.Config,
	) {
		r := chi.NewRouter()

		r.Route("/accounts", func(r chi.Router) {
			r.Post("/", accountController.Post)
			r.Patch("/{ehid}/tenures/{tenureId}/end-date", accountController.PatchEndDate)
		})

		r.Route("/sessions", func(r chi.Router) {
			r.Post("/", sessionController.Post)
		})

		r.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(config.TokenAuth))

			r.Route("/accounts/me/profile", func(r chi.Router) {
				r.Get("/", accountController.GetMyProfile)
			})

			r.Route("/accounts/me/tenures", func(r chi.Router) {
				r.Get("/", accountController.GetMyTenures)
			})
		})

		err := http.ListenAndServe(":8080", r)

		if err != nil {
			panic(err)
		}
	}

	if err := container.Invoke(process); err != nil {
		panic(err)
	}
}
