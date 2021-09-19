package handlers

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"strconv"
	"t_auth/pkg/model"
	"t_auth/pkg/services"
)

func NewLoginHandler(userService *services.UserService, tokenService *services.TokenService) *LoginHandler {
	return &LoginHandler{
		userService:  userService,
		tokenService: tokenService,
	}
}

type LoginHandlerI interface {
	CreateNewUser(w http.ResponseWriter, req *http.Request)
	GetUserProfile(w http.ResponseWriter, req *http.Request)
	Login(w http.ResponseWriter, req *http.Request)
	Refresh(w http.ResponseWriter, req *http.Request)
}

type LoginHandler struct {
	userService  *services.UserService
	tokenService *services.TokenService
}

func (u LoginHandler) GetUserProfile(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		userID := req.Context().Value("CurrentUser").(model.ActiveUserData).ID
		user, err := u.userService.GetUserByID(userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		jUser, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		length, err := w.Write(jUser)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)
	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}

}

func (u LoginHandler) Login(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "POST":
		var loginForm model.LoginRequest
		err := json.NewDecoder(req.Body).Decode(&loginForm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		user, err := u.userService.GetUserByEmail(loginForm.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(loginForm.Password))
		if err != nil {
			http.Error(w, "invalid input", http.StatusUnauthorized)
			return
		}
		accessLifetimeMinutes, _ := strconv.Atoi(os.Getenv("accessLifetimeMinutes"))
		refreshLifetimeMinutes, _ := strconv.Atoi(os.Getenv("refreshLifetimeMinutes"))
		accessString, err := u.tokenService.GenerateToken(user.ID, accessLifetimeMinutes, os.Getenv("accessSecret"))
		refreshString, err := u.tokenService.GenerateToken(user.ID, refreshLifetimeMinutes, os.Getenv("refreshSecret"))
		if err != nil {
			http.Error(w, "Fail to generate tokens", http.StatusUnauthorized)
		}

		resp := &model.TokenPair{
			AccessToken:  accessString,
			RefreshToken: refreshString,
		}
		respJ, _ := json.Marshal(resp)

		w.WriteHeader(http.StatusOK)
		length, err := w.Write(respJ)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

}

func (u LoginHandler) CreateNewUser(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		var user model.User
		err := json.NewDecoder(req.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
		}
		err = u.userService.CreateNewUser(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

	w.WriteHeader(http.StatusOK)

}

func (u LoginHandler) Refresh(w http.ResponseWriter, req *http.Request) {
	userID := req.Context().Value("CurrentUser").(model.ActiveUserData).ID
	accessLifetimeMinutes, _ := strconv.Atoi(os.Getenv("accessLifetimeMinutes"))
	refreshLifetimeMinutes, _ := strconv.Atoi(os.Getenv("refreshLifetimeMinutes"))
	accessString, err := u.tokenService.GenerateToken(userID, accessLifetimeMinutes, os.Getenv("accessSecret"))
	refreshString, err := u.tokenService.GenerateToken(userID, refreshLifetimeMinutes, os.Getenv("refreshSecret"))
	if err != nil {
		http.Error(w, "Fail to generate tokens", http.StatusUnauthorized)
	}

	resp := &model.TokenPair{
		AccessToken:  accessString,
		RefreshToken: refreshString,
	}
	respJ, _ := json.Marshal(resp)

	w.WriteHeader(http.StatusOK)
	length, err := w.Write(respJ)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(length)
}
