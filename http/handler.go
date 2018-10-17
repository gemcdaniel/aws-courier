package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gemcdaniel/aws-courier"
	"github.com/gorilla/mux"
)

// Handler provides the routing for the API using the CredentialsService
type Handler struct {
	*mux.Router

	credentialsService awscourier.CredentialsService
}

// NewHandler initializes a new Handler
func NewHandler() *Handler {
	h := &Handler{
		Router: mux.NewRouter(),
	}

	h.HandleFunc("/aws_keys/{profile}", h.handleCredentials).Methods(http.MethodGet)

	return h
}

// CredentialsService sets the CredentialsService for the Handler
func (h *Handler) CredentialsService(credentialsService awscourier.CredentialsService) {
	h.credentialsService = credentialsService
}

func (h *Handler) handleCredentials(w http.ResponseWriter, r *http.Request) {
	profile, ok := mux.Vars(r)["profile"]
	if !ok {
		writeError(http.StatusBadRequest, fmt.Errorf("no profile provided"), w)
	}

	credentials, err := h.credentialsService.Credentials(&profile)
	if err != nil {
		writeError(http.StatusInternalServerError, err, w)
	}

	data, err := json.Marshal(credentials)
	if err != nil {
		writeError(http.StatusInternalServerError, err, w)
	}

	writeResponse(http.StatusOK, data, w)
}

func writeError(code int, err error, w http.ResponseWriter) {
	b, e := json.Marshal(err)
	if e != nil {
		writeResponse(http.StatusInternalServerError, []byte("{\"Message\":\"failed to marshal error\"}"), w)
	}
	writeResponse(code, b, w)
}

func writeResponse(code int, data []byte, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		w.Write(data)
	}
}
