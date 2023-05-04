package opts

import (
	"fmt"
	"net/http"

	"github.com/go-chi/httplog"
	"github.com/spf13/cobra"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-idp/internal/account"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/credential"
	"github.com/mrexmelle/connect-idp/internal/profile"
	"github.com/mrexmelle/connect-idp/internal/session"
	"github.com/mrexmelle/connect-idp/internal/tenure"
	"go.uber.org/dig"
)

func NewConfig() *config.Config {
	cfg, err := config.New(
		"application", "yaml",
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

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Serve(cmd *cobra.Command, args []string) {
	container := dig.New()
	container.Provide(NewConfig)
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

		r.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://localhost:3000"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
			MaxAge:           300, // Maximum value not ignored by any of major browsers
		}))

		r.Route("/accounts", func(r chi.Router) {
			r.Post("/", accountController.Post)
			r.Post("/{ehid}/tenures", accountController.PostTenure)
			r.Patch("/{ehid}/tenures/{tenureId}/end-date", accountController.PatchEndDate)
		})

		r.Route("/sessions", func(r chi.Router) {
			r.Post("/", sessionController.Post)
		})

		r.Group(func(r chi.Router) {
			logger := httplog.NewLogger("accounts-me-logger", httplog.Options{
				JSON: true,
			})
			r.Use(httplog.RequestLogger(logger))
			r.Use(jwtauth.Verifier(config.TokenAuth))

			r.Route("/accounts/me", func(r chi.Router) {
				r.Get("/profile", accountController.GetMyProfile)
				r.Get("/tenures", accountController.GetMyTenures)
			})
		})

		err := http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r)

		if err != nil {
			panic(err)
		}
	}

	if err := container.Invoke(process); err != nil {
		panic(err)
	}
}

var ServeCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Connect IdP server",
	Run:   Serve,
}
