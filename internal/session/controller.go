package session

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mrexmelle/connect-idp/internal/config"
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

	fmt.Println("Post: Before authenticating")

	queryResult, err := c.SessionService.Authenticate(requestBody)

	if err != nil {
		http.Error(w, "POST failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if !queryResult {
		http.Error(w, "POST failure: "+err.Error(), http.StatusUnauthorized)
		return
	}

	fmt.Println("Post: Before signing")

	signingResult, exp, err := c.SessionService.GenerateJwt(requestBody.EmployeeId)
	if err != nil {
		http.Error(w, "POST failure: "+err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(
		w, &http.Cookie{
			Name:     "jwt",
			Value:    signingResult,
			Expires:  exp,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteNoneMode,
		},
	)
	responseBody, _ := json.Marshal(
		&SessionPostResponse{Token: signingResult},
	)
	w.Write([]byte(responseBody))
}
