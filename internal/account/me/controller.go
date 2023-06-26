package accountMe

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/currentOrganization"
	"github.com/mrexmelle/connect-idp/internal/profile"
	"github.com/mrexmelle/connect-idp/internal/superior"
	"github.com/mrexmelle/connect-idp/internal/tenure"
)

type Controller struct {
	Config                     *config.Config
	CurrentOrganizationService *currentOrganization.Service
	ProfileService             *profile.Service
	SuperiorService            *superior.Service
	TenureService              *tenure.Service
}

func NewController(
	cfg *config.Config,
	cos *currentOrganization.Service,
	ps *profile.Service,
	ss *superior.Service,
	ts *tenure.Service,
) *Controller {
	return &Controller{
		Config:                     cfg,
		CurrentOrganizationService: cos,
		ProfileService:             ps,
		SuperiorService:            ss,
		TenureService:              ts,
	}
}

func (c *Controller) GetTenures(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.TenureService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetProfile(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.ProfileService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetCurrentOrganizations(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.CurrentOrganizationService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetSuperiors(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "GET failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	response := c.SuperiorService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
