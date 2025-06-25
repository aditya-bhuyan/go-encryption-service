package keymanager

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
)

var (
	keyStore = map[string][]byte{}
	mu       sync.Mutex
)

func CreateKey() string {
	mu.Lock()
	defer mu.Unlock()
	key := make([]byte, 32) // AES-256
	rand.Read(key)
	keyID := base64.URLEncoding.EncodeToString(key[:6])
	keyStore[keyID] = key
	return keyID
}

func DeleteKey(id string) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := keyStore[id]; exists {
		delete(keyStore, id)
		return true
	}
	return false
}

func ResetKey(id string) bool {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := keyStore[id]; exists {
		newKey := make([]byte, 32)
		rand.Read(newKey)
		keyStore[id] = newKey
		return true
	}
	return false
}

func GetKey(id string) ([]byte, bool) {
	mu.Lock()
	defer mu.Unlock()
	key, exists := keyStore[id]
	return key, exists
}
