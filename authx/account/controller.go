package account

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
	"github.com/mrexmelle/connect-iam/authx/config"
)

func Post(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody AccountPostRequest
		json.NewDecoder(r.Body).Decode(&requestBody)

		err := Register(requestBody, config.Db)

		if err != nil {
			http.Error(w, "Registration Failure: "+err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody, _ := json.Marshal(
			&AccountPostResponse{Status: "OK"},
		)

		w.Write([]byte(responseBody))
	}
}

func PatchEndDate(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestBody AccountPatchRequest
		json.NewDecoder(r.Body).Decode(&requestBody)

		ehid := chi.URLParam(r, "ehid")
		tenureId, err := strconv.Atoi(chi.URLParam(r, "tenureId"))
		if err != nil {
			http.Error(w, "Patching endDate Failure: "+err.Error(), http.StatusBadRequest)
		}

		err = UpdateEndDate(tenureId, ehid, requestBody, config.Db)
		if err != nil {
			http.Error(w, "Patching endDate Failure: "+err.Error(), http.StatusInternalServerError)
			return
		}

		responseBody, _ := json.Marshal(
			&AccountPatchResponse{Status: "OK"},
		)

		w.Write([]byte(responseBody))
	}
}

func GetMyProfile(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, "Verification Failure: "+err.Error(), http.StatusUnauthorized)
			return
		}
		res, err := RetrieveProfile(claims["sub"].(string), config.Db)

		responseBody, _ := json.Marshal(
			&res,
		)

		w.Write([]byte(responseBody))
	}
}

func GetMyTenures(config *config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, claims, err := jwtauth.FromContext(r.Context())
		if err != nil {
			http.Error(w, "Verification Failure: "+err.Error(), http.StatusUnauthorized)
			return
		}
		res, err := RetrieveTenures(claims["sub"].(string), config.Db)

		responseBody, _ := json.Marshal(
			&res,
		)

		w.Write([]byte(responseBody))
	}
}
