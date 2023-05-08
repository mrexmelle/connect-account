package accountTenure

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-idp/internal/config"
	"github.com/mrexmelle/connect-idp/internal/tenure"
)

type Controller struct {
	Config               *config.Config
	AccountTenureService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:               cfg,
		AccountTenureService: svc,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody tenure.Entity
	json.NewDecoder(r.Body).Decode(&requestBody)

	requestBody.Ehid = chi.URLParam(r, "ehid")
	response := c.AccountTenureService.Create(requestBody)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) PatchEndDate(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	ehid := chi.URLParam(r, "ehid")
	tenureId, err := strconv.Atoi(chi.URLParam(r, "tenureId"))
	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusBadRequest)
	}

	response := c.AccountTenureService.UpdateEndDateByEhidAndTenureId(
		requestBody,
		ehid,
		tenureId,
	)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
