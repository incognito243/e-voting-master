package auth

import (
	"fmt"
	"testing"
)

func TestJWTService(t *testing.T) {
	// Initialize the JWT service
	secretKey := "test_secret_key"
	tokenDuration := 60 * 60 * 24
	jwtService := NewJWTService(secretKey, int64(tokenDuration))

	// Create a test user
	user := &User{
		Username:    "testuser",
		CitizenID:   "123456789",
		CitizenName: "Test User",
		Verified:    true,
		Email:       "testuser@example.com",
		PublicKey:   "test_public_key",
	}

	// Test GenerateToken
	token, err := jwtService.GenerateToken(user)
	if err != nil {
		t.Fatalf("Error generating token: %v", err)
	}
	fmt.Println("Generated Token:", token)

	// Test ValidateToken
	validatedToken, err := jwtService.ValidateToken(token)
	if err != nil {
		t.Fatalf("Error validating token: %v", err)
	}
	fmt.Println("Validated Token:", validatedToken)

	// Test GetClaims
	claims, err := jwtService.GetClaims(token)
	if err != nil {
		t.Fatalf("Error getting claims: %v", err)
	}
	exp, _ := claims.GetExpirationTime()
	fmt.Println("Claims:", exp.Unix())

	userFromClaims, err := jwtService.GetUserFromClaims(token)
	if err != nil {
		t.Fatalf("Error getting user from claims: %v", err)
	}
	fmt.Println("User from Claims:", userFromClaims)
}
