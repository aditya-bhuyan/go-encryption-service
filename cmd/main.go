package main

import (
	"net/http"

	"github.com/aditya-bhuyan/go-encryption-service/internal/api"

	"github.com/aditya-bhuyan/go-encryption-service/pkg/logger"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.RegisterRoutes(r)
	logger.Info("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
