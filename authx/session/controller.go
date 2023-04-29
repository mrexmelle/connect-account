package session

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-iam/authx/config"
)

type Controller struct {
	Config         *config.Config
	SessionService *Service
}

func NewController(cfg *config.Config, svc *Service) *Controller {
	return &Controller{
		Config:         cfg,
		SessionService: svc,
	}
}

func (c *Controller) Post(w http.ResponseWriter, r *http.Request) {
	var requestBody SessionPostRequest
	json.NewDecoder(r.Body).Decode(&requestBody)

	queryResult, err := c.SessionService.Authenticate(requestBody)

	if err != nil {
		http.Error(w, "Authentication Failure", http.StatusInternalServerError)
		return
	}

	if queryResult == false {
		http.Error(w, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	signingResult, err := c.SessionService.GenerateJwt(requestBody.EmployeeId)

	if err != nil {
		http.Error(w, "Signing Failure", http.StatusInternalServerError)
		return
	}

	responseBody, _ := json.Marshal(
		&SessionPostResponse{Token: signingResult},
	)

	w.Write([]byte(responseBody))
}
