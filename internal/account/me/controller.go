package accountMe

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-idp/internal/accountOrganization"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/credential"
	"github.com/mrexmelle/connect-idp/internal/profile"
	"github.com/mrexmelle/connect-idp/internal/superior"
	"github.com/mrexmelle/connect-idp/internal/tenure"
)

type Controller struct {
	Config                     *config.Config
	AccountOrganizationService *accountOrganization.Service
	CredentialService          *credential.Service
	ProfileService             *profile.Service
	SuperiorService            *superior.Service
	TenureService              *tenure.Service
}

func NewController(
	cfg *config.Config,
	aos *accountOrganization.Service,
	cs *credential.Service,
	ps *profile.Service,
	ss *superior.Service,
	ts *tenure.Service,
) *Controller {
	return &Controller{
		Config:                     cfg,
		AccountOrganizationService: aos,
		CredentialService:          cs,
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
	response := c.SuperiorService.RetrieveByEhid(claims["sub"].(string))
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) PatchPassword(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusUnauthorized)
		return
	}

	var requestBody credential.PatchRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if requestBody.CurrentPassword == requestBody.NewPassword {
		responseBody, _ := json.Marshal(&credential.ResponseDto{
			Status: "OK",
		})
		w.Write([]byte(responseBody))
		return
	}

	profileResponse := c.ProfileService.RetrieveByEhid(claims["sub"].(string))
	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusInternalServerError)
		return
	}
	requestBody.EmployeeId = profileResponse.Profile.EmployeeId
	response, err := c.CredentialService.PatchPassword(requestBody)
	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusNotFound)
		return
	}
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
