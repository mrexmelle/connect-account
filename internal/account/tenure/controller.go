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
	Config        *config.Config
	TenureService *tenure.Service
}

func NewController(cfg *config.Config, ts *tenure.Service) *Controller {
	return &Controller{
		Config:        cfg,
		TenureService: ts,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody tenure.Entity
	json.NewDecoder(r.Body).Decode(&requestBody)

	requestBody.Ehid = chi.URLParam(r, "ehid")
	response := c.TenureService.Create(requestBody)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) PatchEndDate(w http.ResponseWriter, r *http.Request) {
	var requestBody tenure.PatchRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	ehid := chi.URLParam(r, "ehid")
	tenureId, err := strconv.Atoi(chi.URLParam(r, "tenureId"))
	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusBadRequest)
	}

	response := c.TenureService.UpdateEndDateByEhidAndTenureId(
		requestBody,
		ehid,
		tenureId,
	)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
