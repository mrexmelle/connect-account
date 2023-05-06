package account

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-idp/internal/config"
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
		&AccountPostResponse{Status: "OK"},
	)
	w.Write([]byte(responseBody))
}
