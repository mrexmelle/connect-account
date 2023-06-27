package tenure

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-idp/internal/config"
)

type Controller struct {
	Config        *config.Config
	TenureService *Service
}

func NewController(cfg *config.Config, ts *Service) *Controller {
	return &Controller{
		Config:        cfg,
		TenureService: ts,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody Entity
	json.NewDecoder(r.Body).Decode(&requestBody)

	response := c.TenureService.Create(requestBody)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) PatchEndDate(w http.ResponseWriter, r *http.Request) {
	var requestBody PatchRequestDto
	json.NewDecoder(r.Body).Decode(&requestBody)

	tenureId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "PATCH failure: "+err.Error(), http.StatusBadRequest)
	}

	response := c.TenureService.UpdateEndDateById(
		requestBody,
		tenureId,
	)
	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
