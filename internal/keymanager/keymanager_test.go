package keymanager

import (
	"testing"
)

func TestKeyLifecycle(t *testing.T) {
	keyID := CreateKey()
	if keyID == "" {
		t.Fatal("expected a non-empty key ID")
	}

	key, exists := GetKey(keyID)
	if !exists || len(key) != 32 {
		t.Fatal("expected to retrieve a 32-byte key")
	}

	if !ResetKey(keyID) {
		t.Fatal("expected key to be reset")
	}

	if !DeleteKey(keyID) {
		t.Fatal("expected key to be deleted")
	}

	_, exists = GetKey(keyID)
	if exists {
		t.Fatal("expected key to be absent after deletion")
	}
}
