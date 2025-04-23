package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type IJWTService interface {
	GenerateToken(user *User) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
	GetClaims(token string) (jwt.MapClaims, error)
	GetUserFromClaims(encodedToken string) (*User, error)
}
