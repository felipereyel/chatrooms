package utils

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	id := "123"
	user := "test"

	token, expiration, err := GenerateJWT(id, user)
	if err != nil {
		t.Fatalf("Error generating JWT: %v", err)
	}

	if token == "" {
		t.Fatalf("Token is empty")
	}

	if expiration.Before(time.Now()) {
		t.Fatalf("Expiration is before current time")
	}

	idFromToken, err := ParseJWT(token)
	if err != nil {
		t.Fatalf("Error parsing JWT: %v", err)
	}

	if idFromToken != id {
		t.Fatalf("Expected %v, got %v", id, idFromToken)
	}
}
