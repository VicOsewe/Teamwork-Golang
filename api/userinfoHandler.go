package api

import (
	"Teamwork-Golang/getting"
	"Teamwork-Golang/registering"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreateUser(service registering.RegisterService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := registering.Users{}
		// read the request body into byte slice
		requestBody, _ := ioutil.ReadAll(r.Body)
		// parse the body into the user slice, handle any errors
		if err := json.Unmarshal(requestBody, &user); err != nil {
			writeError(w, err, string(requestBody))
			return
		}
		// send user slice to the registering service, handle any errors
		UserID, err := service.CreateUser(user)
		if err != nil {
			writeError(w, err, user)
			return
		}

		// write a success response back to the originating service
		var response = operationResponse{"SUCCESS", "user account successfully", UserID, make([]string, 0)}
		writeJSON(w, response)
	}
}

func UserSignIn(service getting.GettingService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userinfo := getting.UserSignInfo{}
		err := json.NewDecoder(r.Body).Decode(&userinfo)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		erro := service.UserSignIn(userinfo)
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
}
