package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-iam/idp/account"
	"github.com/mrexmelle/connect-iam/idp/config"
	"github.com/mrexmelle/connect-iam/idp/credential"
	"github.com/mrexmelle/connect-iam/idp/profile"
	"github.com/mrexmelle/connect-iam/idp/session"
	"github.com/mrexmelle/connect-iam/idp/tenure"
	"go.uber.org/dig"
)

func Config() *config.Config {
	cfg, err := config.New(
		"idp", "yaml",
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
	container.Provide(credential.NewRepository)
	container.Provide(profile.NewRepository)
	container.Provide(tenure.NewRepository)
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

			r.Route("/accounts/me", func(r chi.Router) {
				r.Get("/profile", accountController.GetMyProfile)
				r.Get("/tenures", accountController.GetMyTenures)
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
