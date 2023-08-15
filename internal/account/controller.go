package account

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-idp/internal/accountOrganization"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/credential"
	"github.com/mrexmelle/connect-idp/internal/mapper"
	"github.com/mrexmelle/connect-idp/internal/profile"
	"github.com/mrexmelle/connect-idp/internal/superior"
	"github.com/mrexmelle/connect-idp/internal/tenure"
)

type Controller struct {
	Config                     *config.Config
	AccountService             *Service
	AccountOrganizationService *accountOrganization.Service
	CredentialService          *credential.Service
	ProfileService             *profile.Service
	SuperiorService            *superior.Service
	TenureService              *tenure.Service
}

func NewController(
	cfg *config.Config,
	as *Service,
	aos *accountOrganization.Service,
	cs *credential.Service,
	ps *profile.Service,
	ss *superior.Service,
	ts *tenure.Service,
) *Controller {
	return &Controller{
		Config:                     cfg,
		AccountService:             as,
		AccountOrganizationService: aos,
		CredentialService:          cs,
		ProfileService:             ps,
		SuperiorService:            ss,
		TenureService:              ts,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody AccountPostRequest
	json.NewDecoder(r.Body).Decode(&requestBody)

	err := c.AccountService.Register(requestBody)

	if err != nil {
		http.Error(w, "POST failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&AccountResponse{Status: "OK"},
	)
	w.Write([]byte(responseBody))
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	employeeId := chi.URLParam(r, "employee_id")

	err := c.AccountService.DeleteByEmployeeId(employeeId)
	if err != nil {
		http.Error(w, "DELETE failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	ehid := mapper.ToEhid(employeeId)
	err = c.AccountService.DeleteEmailByEhid(ehid)
	if err != nil {
		http.Error(w, "DELETE failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&AccountResponse{Status: "OK"},
	)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetTenures(w http.ResponseWriter, r *http.Request) {
	ehid := chi.URLParam(r, "ehid")

	response := c.TenureService.RetrieveByEhid(ehid)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetProfile(w http.ResponseWriter, r *http.Request) {
	ehid := chi.URLParam(r, "ehid")

	response := c.ProfileService.RetrieveByEhid(ehid)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetOrganizations(w http.ResponseWriter, r *http.Request) {
	ehid := chi.URLParam(r, "ehid")

	response := c.AccountOrganizationService.RetrieveByEhid(ehid)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetSuperiors(w http.ResponseWriter, r *http.Request) {
	ehid := chi.URLParam(r, "ehid")

	response := c.SuperiorService.RetrieveByEhid(ehid)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) PostPasswordReset(w http.ResponseWriter, r *http.Request) {
	ehid := chi.URLParam(r, "ehid")

	profileResponse := c.ProfileService.RetrieveByEhid(ehid)
	err := c.CredentialService.ResetPassword(profileResponse.Profile.EmployeeId)
	if err != nil {
		http.Error(w, "POST failure: "+err.Error(), http.StatusNotFound)
		return
	}
	responseBody, _ := json.Marshal(&credential.ResponseDto{Status: "OK"})
	w.Write([]byte(responseBody))
}
