package session

import (
	"encoding/json"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var requestBody SessionRequest
	json.NewDecoder(r.Body).Decode(&requestBody)

	queryResult, err := Authenticate(requestBody.Id, requestBody.Password)

	if err != nil {
		http.Error(w, "Authentication Failure", http.StatusInternalServerError)
		return
	}

	if queryResult == false {
		http.Error(w, "Authentication Failure", http.StatusUnauthorized)
		return
	}

	signingResult, err := GenerateJwt(requestBody.Id)

	if err != nil {
		http.Error(w, "Signing Failure", http.StatusInternalServerError)
		panic(err)
		return
	}

	sessionResponse := SessionResponse{Token: signingResult}

	responseBody, _ := json.Marshal(&sessionResponse)

	w.Write([]byte(responseBody))
}
