package services

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"strings"
	"time"
)

type TokenServiceI interface {
	GenerateToken(userID, lifetimeMinutes int, secret string) (string, error)
	ValidateToken(tokenString, secret string) (*JwtCustomClaims, error)
	GetTokenFromBearerString(bearerString string) string
	CheckUID(uID string) (int, error)
}
type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.StandardClaims
	UID string `json:"u_id"`
}

func NewTokenService() *TokenService {
	return &TokenService{}
}

type TokenService struct {
}

func (t TokenService) GenerateToken(userID, lifetimeMinutes int, secret string) (string, error) {
	uID := uuid.New().String()
	claims := &JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(lifetimeMinutes)).Unix(),
		},
		uID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func (t TokenService) ValidateToken(tokenString, secret string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

func (t TokenService) GetTokenFromBearerString(bearerString string) string {
	if bearerString == "" {
		return ""
	}

	parts := strings.Split(bearerString, "Bearer")
	if len(parts) != 2 {
		return ""
	}

	token := strings.TrimSpace(parts[1])
	if len(token) < 1 {
		return ""
	}

	return token
}
func (t TokenService) CheckUID(uID string) (int, error) {
	userID, err := t.tokenRepository.GetByUID(uID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
