package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	// Register all API endpoints
	r.HandleFunc("/api/key", CreateKey).Methods("POST")
	r.HandleFunc("/api/key/{id}", DeleteKey).Methods("DELETE")
	r.HandleFunc("/api/key/{id}/reset", ResetKey).Methods("PUT")
	r.HandleFunc("/api/encrypt", EncryptData).Methods("POST")
	r.HandleFunc("/api/decrypt", DecryptData).Methods("POST")
	// inside RegisterRoutes
	r.PathPrefix("/docs/").Handler(http.StripPrefix("/docs/", http.FileServer(http.Dir("./internal/docs"))))
}
