package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-encryption-service/internal/keymanager"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()
	RegisterRoutes(r)
	return r
}

func TestCreateKey(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("POST", "/api/key", nil)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", resp.Code)
	}
}

func TestEncryptDecryptCycle(t *testing.T) {
	r := setupRouter()
	keyID := keymanager.CreateKey()

	payload := map[string]string{
		"key_id": keyID,
		"data":   "hello",
	}
	jsonPayload, _ := json.Marshal(payload)

	// Encrypt
	req, _ := http.NewRequest("POST", "/api/encrypt", bytes.NewReader(jsonPayload))
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("encrypt failed with status %d", resp.Code)
	}

	var encResp map[string]string
	json.NewDecoder(resp.Body).Decode(&encResp)
	encData := encResp["encrypted"]

	// Decrypt
	payload["data"] = encData
	jsonPayload, _ = json.Marshal(payload)
	req, _ = http.NewRequest("POST", "/api/decrypt", bytes.NewReader(jsonPayload))
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)
	if resp.Code != http.StatusOK {
		t.Fatalf("decrypt failed with status %d", resp.Code)
	}

	var decResp map[string]string
	json.NewDecoder(resp.Body).Decode(&decResp)
	if decResp["decrypted"] != "hello" {
		t.Fatalf("expected 'hello', got '%s'", decResp["decrypted"])
	}
}
