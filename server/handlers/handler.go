package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/RemcoE33/share-secret/store"
)

// inject dependencies here
type Server struct {
	Storer store.Storer
}

// init new server with the dependencies
func NewServer(store store.Storer) *Server {
	return &Server{
		Storer: store,
	}
}

// small helper function to serve a 200 response to the client
func serveJSON(payload any, w http.ResponseWriter) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(payload)
	return err
}

// returns a not found to the client
func serveNotFoundOrPastDeadline(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "secret is not found or exired",
	})
}

// internal server error that logs the error but not serve to the client
func serveInternalError(err error, w http.ResponseWriter) {
	fmt.Println(err)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "something went wrong",
	})
}
