package organizationTree

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mrexmelle/connect-idp/internal/config"
)

type Controller struct {
	Config                  *config.Config
	OrganizationTreeService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:                  cfg,
		OrganizationTreeService: svc,
	}
}

func (c *Controller) GetSiblingsAndAncestralSiblings(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	response := c.OrganizationTreeService.RetrieveSiblingsAndAncestralSiblingsById(id)
	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}

func (c *Controller) GetChildren(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	response := c.OrganizationTreeService.RetrieveChildrenById(id)
	if response.Status != "OK" {
		http.Error(w, "GET failure: "+response.Status, http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(&response)
	w.Write([]byte(responseBody))
}
