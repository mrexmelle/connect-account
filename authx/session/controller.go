package session

import (
	"encoding/json"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	var requestBody SessionRequest
	json.NewDecoder(r.Body).Decode(&requestBody)

	queryResult := Authenticate(requestBody.Id, requestBody.Password)

	var sessionResponse SessionResponse
	if queryResult {
		sessionResponse = SessionResponse{Token: "valid"}
	} else {
		sessionResponse = SessionResponse{Token: "invalid"}
	}

	responseBody, _ := json.Marshal(&sessionResponse)

	w.Write([]byte(responseBody))
}
