package encryption

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	key := make([]byte, 32) // AES-256
	rand.Read(key)
	plaintext := []byte("Test Message")

	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	decrypted, err := Decrypt(ciphertext, key)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	if !bytes.Equal(decrypted, plaintext) {
		t.Errorf("Expected %s, got %s", plaintext, decrypted)
	}
}
