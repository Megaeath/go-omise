package decrypt

import (
	"testing"
)

func TestROT128(t *testing.T) {
	input := []byte("hello")
	encrypted := ROT128(input)
	decrypted := ROT128(encrypted)

	if string(decrypted) != "hello" {
		t.Errorf("Expected %s, got %s", "hello", decrypted)
	}
}
