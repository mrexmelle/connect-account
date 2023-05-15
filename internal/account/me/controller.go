package accountMe

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	accountOrganization "github.com/mrexmelle/connect-idp/internal/account/organization"
	accountProfile "github.com/mrexmelle/connect-idp/internal/account/profile"
	accountSuperior "github.com/mrexmelle/connect-idp/internal/account/superior"
	accountTenure "github.com/mrexmelle/connect-idp/internal/account/tenure"
	"github.com/mrexmelle/connect-idp/internal/config"
)

type Controller struct {
	Config                     *config.Config
	AccountOrganizationService *accountOrganization.Service
	AccountProfileService      *accountProfile.Service
	AccountSuperiorService     *accountSuperior.Service
	AccountTenureService       *accountTenure.Service
}

func NewController(
	cfg *config.Config,
	aos *accountOrganization.Service,
	aps *accountProfile.Service,
	ass *accountSuperior.Service,
	ats *accountTenure.Service,
) *Controller {
	return &Controller{
		Config:                     cfg,
		AccountOrganizationService: aos,
		AccountProfileService:      aps,
		AccountSuperiorService:     ass,
		AccountTenureService:       ats,
	}
}

func (c *Controller) GetTenures(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.AccountTenureService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetProfile(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.AccountProfileService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetOrganizations(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.AccountOrganizationService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetSuperiors(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.AccountSuperiorService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
