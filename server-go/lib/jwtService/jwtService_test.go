package jwtService

import (
	"server-go/lib/databaseService"
	"testing"
)

func TestGenerateJWT(t *testing.T) {
	// Test the GenerateJWT function
	user := &databaseService.User{
		UserID: "test",
	}
	_, err := GenerateJWT(user)
	if err != nil {
		t.Errorf("GenerateJWT(%v) = %v; want nil", user, err)
	}
}

func TestParseToken(t *testing.T) {
	user := &databaseService.User{
		UserID: "test",
	}
	token, _ := GenerateJWT(user)
	_, err := ParseToken(*token)
	if err != nil {
		t.Errorf("ParseToken(%s) = %v; want nil", *token, err)
	}
}

func TestParseToken2(t *testing.T) {
	token := "invalid_token"
	_, err := ParseToken(token)
	if err == nil {
		t.Errorf("ParseToken(%s) = %v; want error", token, err)
	}
}

func TestExtractClaims(t *testing.T) {
	token := "invalid_token"

	_, err := ParseToken(token)

	if err == nil {
		t.Errorf("ParseToken(%s) = %v; want error", token, err)
	}
}
