package account

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/mapper"
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
