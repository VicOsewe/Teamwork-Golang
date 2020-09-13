package api

import (
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
		// parse the body into the sites slice, handle any errors
		if err := json.Unmarshal(requestBody, &user); err != nil {
			writeError(w, err, string(requestBody))
			return
		}
		// send sites slice to the registering service, handle any errors
		if err := service.CreateUser(user); err != nil {
			writeError(w, err, user)
			return
		}
		// write a success response back to the originating service
		var response = operationResponse{"SUCCESS", "sites registered to database", make([]string, 0)}
		writeJSON(w, response)
	}
}
