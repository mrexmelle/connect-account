package session

import (
	"encoding/json"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var requestBody SessionPostRequest
	json.NewDecoder(r.Body).Decode(&requestBody)

	queryResult, err := Authenticate(requestBody)

	if err != nil {
		http.Error(w, "Authentication Failure", http.StatusInternalServerError)
		return
	}

	if queryResult == false {
		http.Error(w, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	signingResult, err := GenerateJwt(requestBody.EmployeeId)

	if err != nil {
		http.Error(w, "Signing Failure", http.StatusInternalServerError)
		panic(err)
		return
	}

	responseBody, _ := json.Marshal(
		&SessionPostResponse{Token: signingResult},
	)

	w.Write([]byte(responseBody))
}
