package api

import (
	"Teamwork-Golang/registering"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/muhoro/log"
)

type operationResponse struct {
	Status  string
	Message string
	UserID  uuid.UUID
	Errors  []string
}

type articleResponse struct {
	Status       string
	Message      string
	ArticleId    uuid.UUID
	CreatedOn    time.Time
	ArticleTitle string
	Errors       []string
}

func writeJSON(w http.ResponseWriter, resp interface{}) {
	// set the appropriate headers
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// pack the response struct into a JSON response body and send it
	responseBytes, _ := json.Marshal(&resp)
	w.Write(responseBytes)
}

func writeError(w http.ResponseWriter, err error, data interface{}) {
	// set the appropriate headers
	w.Header().Add("Content-Type", "application/json")
	resp := operationResponse{
		Status: "FAILED",
	}

	// depending on the type of the error...
	switch err.(type) {
	// if it is a custom error from the registering service,
	// set the status 400 with the custom errors in the response body...
	case *registering.RegisteringError:
		w.WriteHeader(http.StatusBadRequest)
		resp.Errors = strings.Split(err.Error(), ",")
	// otherwise, send back a generic internal server error 500 with new error
	default:
		log.Error(err.Error(), data)
		resp.Errors = []string{err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
	}

	// pack the response struct into a JSON response body and send it
	responseBytes, _ := json.Marshal(&resp)
	w.Write(responseBytes)
}
