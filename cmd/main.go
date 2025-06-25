package main

import (
	"log"
	"net/http"

	"go-encryption-service/internal/api"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	api.RegisterRoutes(r)
	log.Println("Server starting on :8080")
	http.ListenAndServe(":8080", r)
}
