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
	accountMe "github.com/mrexmelle/connect-idp/internal/account/me"
	accountProfile "github.com/mrexmelle/connect-idp/internal/account/profile"
	accountTenure "github.com/mrexmelle/connect-idp/internal/account/tenure"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/credential"
	"github.com/mrexmelle/connect-idp/internal/organization"
	organizationMember "github.com/mrexmelle/connect-idp/internal/organization/member"
	organizationTree "github.com/mrexmelle/connect-idp/internal/organization/tree"
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
	container.Provide(organization.NewRepository)
	container.Provide(organizationMember.NewRepository)

	container.Provide(account.NewService)
	container.Provide(accountProfile.NewService)
	container.Provide(accountTenure.NewService)
	container.Provide(session.NewService)
	container.Provide(organization.NewService)
	container.Provide(organizationTree.NewService)
	container.Provide(organizationMember.NewService)

	container.Provide(account.NewController)
	container.Provide(accountTenure.NewController)
	container.Provide(accountMe.NewController)
	container.Provide(session.NewController)
	container.Provide(organization.NewController)
	container.Provide(organizationTree.NewController)
	container.Provide(organizationMember.NewController)

	process := func(
		accountController *account.Controller,
		accountTenureController *accountTenure.Controller,
		accountMeController *accountMe.Controller,
		organizationController *organization.Controller,
		organizationMemberController *organizationMember.Controller,
		organizationTreeController *organizationTree.Controller,
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
			r.Delete("/{employee_id}", accountController.Delete)
		})

		r.Route("/accounts/{ehid}/tenures", func(r chi.Router) {
			r.Post("/", accountTenureController.Post)
			r.Patch("/{tenureId}/end-date", accountTenureController.PatchEndDate)
		})

		r.Route("/sessions", func(r chi.Router) {
			r.Post("/", sessionController.Post)
		})

		r.Route("/organizations", func(r chi.Router) {
			r.Post("/", organizationController.Post)
			r.Get("/{id}", organizationController.Get)
			r.Delete("/{id}", organizationController.Delete)
		})

		r.Route("/organizations/{id}/members", func(r chi.Router) {
			r.Get("/", organizationMemberController.GetByOrganizationId)
		})

		r.Route("/organizations/{id}/siblings-and-ancestral-siblings", func(r chi.Router) {
			r.Get("/", organizationTreeController.GetSiblingsAndAncestralSiblings)
		})

		r.Route("/organizations/{id}/children", func(r chi.Router) {
			r.Get("/", organizationTreeController.GetChildren)
		})

		r.Route("/organizations/{id}/lineage", func(r chi.Router) {
			r.Get("/", organizationTreeController.GetLineage)
		})

		r.Group(func(r chi.Router) {
			logger := httplog.NewLogger("secure-path-logger", httplog.Options{
				JSON: true,
			})
			r.Use(httplog.RequestLogger(logger))
			r.Use(jwtauth.Verifier(config.TokenAuth))

			r.Route("/accounts/me", func(r chi.Router) {
				r.Get("/profile", accountMeController.GetProfile)
				r.Get("/tenures", accountMeController.GetTenures)
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
