package session

import (
	"encoding/json"
	"net/http"

	"github.com/mrexmelle/connect-iam/authx/config"
)

func Post(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody SessionPostRequest
		json.NewDecoder(r.Body).Decode(&requestBody)

		queryResult, err := Authenticate(requestBody, config.Db)

		if err != nil {
			http.Error(w, "Authentication Failure", http.StatusInternalServerError)
			return
		}

		if queryResult == false {
			http.Error(w, "Authentication Failure", http.StatusUnauthorized)
			return
		}

		signingResult, err := GenerateJwt(requestBody.EmployeeId, config.TokenAuth)

		if err != nil {
			http.Error(w, "Signing Failure", http.StatusInternalServerError)
			return
		}

		responseBody, _ := json.Marshal(
			&SessionPostResponse{Token: signingResult},
		)

		w.Write([]byte(responseBody))
	}
}
