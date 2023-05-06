package accountMe

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	accountProfile "github.com/mrexmelle/connect-idp/internal/account/profile"
	accountTenure "github.com/mrexmelle/connect-idp/internal/account/tenure"
	"github.com/mrexmelle/connect-idp/internal/config"
)

type Controller struct {
	Config                *config.Config
	AccountTenureService  *accountTenure.Service
	AccountProfileService *accountProfile.Service
}

func NewController(
	cfg *config.Config,
	ats *accountTenure.Service,
	aps *accountProfile.Service,
) *Controller {
	return &Controller{
		Config:                cfg,
		AccountTenureService:  ats,
		AccountProfileService: aps,
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
