package organization

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-idp/internal/config"
)

type Controller struct {
	Config              *config.Config
	OrganizationService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:              cfg,
		OrganizationService: svc,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody Entity
	json.NewDecoder(r.Body).Decode(&requestBody)

	response := c.OrganizationService.Create(requestBody)
	if response.Status != "OK" {
		http.Error(w, "POST failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	response := c.OrganizationService.RetrieveById(id)
	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	response := c.OrganizationService.DeleteById(id)
	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
