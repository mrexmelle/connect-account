package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-iam/idp/config"
)

type Controller struct {
	Config         *config.Config
	AccountService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:         cfg,
		AccountService: svc,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody AccountPostRequest
	json.NewDecoder(r.Body).Decode(&requestBody)

	err := c.AccountService.Register(requestBody)

	if err != nil {
		http.Error(w, "Registration Failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&AccountPostResponse{Status: "OK"},
	)

	w.Write([]byte(responseBody))
}

func (c *Controller) PatchEndDate(w http.ResponseWriter, r *http.Request) {
	var requestBody AccountPatchRequest
	json.NewDecoder(r.Body).Decode(&requestBody)

	ehid := chi.URLParam(r, "ehid")
	tenureId, err := strconv.Atoi(chi.URLParam(r, "tenureId"))
	if err != nil {
		http.Error(w, "Patching endDate Failure: "+err.Error(), http.StatusBadRequest)
	}

	err = c.AccountService.UpdateEndDate(ehid, tenureId, requestBody)
	if err != nil {
		http.Error(w, "Patching endDate Failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&AccountPatchResponse{Status: "OK"},
	)

	w.Write([]byte(responseBody))
}

func (c *Controller) GetMyProfile(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "Verification Failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	res, err := c.AccountService.RetrieveProfile(claims["sub"].(string))

	responseBody, _ := json.Marshal(&res)

	w.Write([]byte(responseBody))
}

func (c *Controller) GetMyTenures(w http.ResponseWriter, r *http.Request) {
	_, claims, err := jwtauth.FromContext(r.Context())
	if err != nil {
		http.Error(w, "Verification Failure: "+err.Error(), http.StatusUnauthorized)
		return
	}
	res, err := c.AccountService.RetrieveTenures(claims["sub"].(string))

	responseBody, _ := json.Marshal(&res)

	w.Write([]byte(responseBody))
}
