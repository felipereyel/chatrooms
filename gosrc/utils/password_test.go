package utils

import "testing"

func TestPassord(t *testing.T) {
	password := "secret"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	// security check
	if password == hash {
		t.Fatalf("Hash and password are the same")
	}

	if !CheckPasswordHash(password, hash) {
		t.Fatalf("Hash and password don't match")
	}
}
