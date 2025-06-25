package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type DataPayload struct {
	KeyID string `json:"key_id"`
	Data  string `json:"data"`
}

func main() {
	// Create a new key
	keyID := createKey()

	// Encrypt data
	cipherText := encryptData(keyID, "Hello Golang!")

	// Decrypt data
	decryptData(keyID, cipherText)

	// Reset the key
	resetKey(keyID)

	// Delete the key
	deleteKey(keyID)
}

func createKey() string {
	resp, err := http.Post("http://localhost:8080/api/key", "application/json", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result map[string]string
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Created Key ID:", result["key_id"])
	return result["key_id"]
}

func encryptData(keyID, data string) string {
	payload := DataPayload{KeyID: keyID, Data: data}
	body, _ := json.Marshal(payload)
	resp, err := http.Post("http://localhost:8080/api/encrypt", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result := make(map[string]string)
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Encrypted:", result["encrypted"])
	return result["encrypted"]
}

func decryptData(keyID, encrypted string) {
	payload := DataPayload{KeyID: keyID, Data: encrypted}
	body, _ := json.Marshal(payload)
	resp, err := http.Post("http://localhost:8080/api/decrypt", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result := make(map[string]string)
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Decrypted:", result["decrypted"])
}

func resetKey(keyID string) {
	req, _ := http.NewRequest("PUT", fmt.Sprintf("http://localhost:8080/api/key/%s/reset", keyID), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Reset Key Response:", string(body))
}

func deleteKey(keyID string) {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("http://localhost:8080/api/key/%s", keyID), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("Delete Key Response:", string(body))
}
