package api

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-encryption-service/internal/encryption"
	"go-encryption-service/internal/keymanager"
	"go-encryption-service/pkg/logger"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type DataPayload struct {
	KeyID string `json:"key_id"`
	Data  string `json:"data"`
}

func CreateKey(w http.ResponseWriter, r *http.Request) {
	keyID := keymanager.CreateKey()
	logger.Info("Key created: " + keyID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"key_id": keyID})
}

func DeleteKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyID := vars["id"]
	ok := keymanager.DeleteKey(keyID)
	if !ok {
		logger.Error("Delete failed: Key " + keyID + " not found")
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	logger.Info("Key deleted: " + keyID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key %s deleted", keyID)
}

func ResetKey(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	keyID := vars["id"]
	ok := keymanager.ResetKey(keyID)
	if !ok {
		logger.Error("Reset failed: Key " + keyID + " not found")
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	logger.Info("Key reset: " + keyID)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key %s reset", keyID)
}

func EncryptData(w http.ResponseWriter, r *http.Request) {
	var payload DataPayload
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &payload)
	key, exists := keymanager.GetKey(payload.KeyID)
	if !exists {
		logger.Error("Encrypt failed: Key " + payload.KeyID + " not found")
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	encData, err := encryption.Encrypt([]byte(payload.Data), key)
	if err != nil {
		logger.Error("Encryption failed: " + err.Error())
		http.Error(w, "Encryption failed", http.StatusInternalServerError)
		return
	}
	encoded := base64.StdEncoding.EncodeToString(encData)
	logger.Info("Data encrypted with key: " + payload.KeyID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"encrypted": encoded})
}

func DecryptData(w http.ResponseWriter, r *http.Request) {
	var payload DataPayload
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &payload)
	key, exists := keymanager.GetKey(payload.KeyID)
	if !exists {
		logger.Error("Decrypt failed: Key " + payload.KeyID + " not found")
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}
	decoded, err := base64.StdEncoding.DecodeString(payload.Data)
	if err != nil {
		logger.Error("Base64 decode failed: " + err.Error())
		http.Error(w, "Invalid base64 data", http.StatusBadRequest)
		return
	}
	decData, err := encryption.Decrypt(decoded, key)
	if err != nil {
		logger.Error("Decryption failed: " + err.Error())
		http.Error(w, "Decryption failed", http.StatusInternalServerError)
		return
	}
	logger.Info("Data decrypted with key: " + payload.KeyID)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"decrypted": string(decData)})
}
