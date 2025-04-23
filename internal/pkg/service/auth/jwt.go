package auth

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

var service *JwtService

func NewJWTService(secretKey string, expire int64) IJWTService {
	if service == nil {
		service = &JwtService{
			secretKey:     secretKey,
			tokenDuration: time.Duration(expire) * time.Second,
		}
	}
	return service
}

func Instance() IJWTService {
	return service
}

func (s *JwtService) GenerateToken(user *User) (string, error) {
	claims := &jwtCustomClaims{
		User: user,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.tokenDuration)),
			Subject:   user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *JwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot parse claims")
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, errors.New("cannot parse expiration time")
	}

	if exp.Before(time.Now()) {
		return nil, errors.New("token has expired")
	}

	return token, nil
}

func (s *JwtService) GetClaims(encodedToken string) (jwt.MapClaims, error) {
	token, err := s.ValidateToken(encodedToken)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot parse claims")
	}
	return claims, nil
}

func (s *JwtService) GetUserFromClaims(encodedToken string) (*User, error) {
	token, err := s.ValidateToken(encodedToken)
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot parse claims")
	}

	userJson, ok := claims["user"]
	if !ok {
		return nil, errors.New("user not found in claims")
	}

	jsonData, err := json.Marshal(userJson)
	if err != nil {
		return nil, err
	}

	var user *User

	if err = json.Unmarshal(jsonData, &user); err != nil {
		return nil, err
	}

	return user, nil
}
