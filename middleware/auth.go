package middleware

import (
	"context"
	"net/http"
	"nix_education/model"
	"nix_education/services"
	"os"
)

func NewAuthMiddlware(tokenService services.TokenServiceI) *AuthMiddlware {
	return &AuthMiddlware{
		tokenService: tokenService,
	}
}

type AuthMiddlwareI interface {
	AccessTokenCheck(next http.HandlerFunc) http.HandlerFunc
	RefreshTokenCheck(next http.HandlerFunc) http.HandlerFunc
}

type AuthMiddlware struct {
	tokenService services.TokenServiceI
}

func (m AuthMiddlware) AccessTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		bearerString := req.Header.Get("Authorization")
		tokenString := m.tokenService.GetTokenFromBearerString(bearerString)
		claims, err := m.tokenService.ValidateToken(tokenString, os.Getenv("accessSecret"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userDB, err := m.tokenService.CheckUID(claims.UID)
		if claims.ID != userDB {
			http.Error(w, "logout", http.StatusUnauthorized)
			return
		}
		currentUser := model.CurrentUser{
			ID: claims.ID,
		}
		req = req.WithContext(context.WithValue(req.Context(), "CurrentUser", currentUser))
		next(w, req)
	}
}

func (m AuthMiddlware) RefreshTokenCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		bearerString := req.Header.Get("Authorization")
		tokenString := m.tokenService.GetTokenFromBearerString(bearerString)
		claims, err := m.tokenService.ValidateToken(tokenString, os.Getenv("refreshSecret"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		userDB, err := m.tokenService.CheckUID(claims.UID)
		if claims.ID != userDB {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		currentUser := model.CurrentUser{
			ID: claims.ID,
		}
		req = req.WithContext(context.WithValue(req.Context(), "CurrentUser", currentUser))
		next(w, req)
	}
}
