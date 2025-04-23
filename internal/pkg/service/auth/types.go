package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	User *User `json:"user"`
	jwt.RegisteredClaims
}

type User struct {
	Username    string `json:"username"`
	CitizenID   string `json:"citizen_id"`
	CitizenName string `json:"citizen_name"`
	Verified    bool   `json:"verified"`
	Email       string `json:"email"`
	PublicKey   string `json:"public_key"`
}
