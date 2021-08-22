package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"nix_education/pkg/model"
	"nix_education/pkg/model/DBrepositories"
	"strconv"
)

type UserHandler struct {
	userRepository DBrepositories.UserDBRepositoryI
}

func NewUserHandler(userRepo DBrepositories.UserDBRepositoryI) *UserHandler {
	return &UserHandler{userRepository: userRepo}
}

func (uh UserHandler) InitHandleFuncRoutes(router *mux.Router) {

	router.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		err = uh.userRepository.CreateUser(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	}).Methods(http.MethodPost)

	router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		users, err := uh.userRepository.GetAllUsers()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(users)
	}).Methods(http.MethodGet)

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		strId := mux.Vars(r)["id"]
		id, err := strconv.Atoi(strId)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		user, err := uh.userRepository.GetUser(int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}).Methods(http.MethodGet)

	router.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		strId := mux.Vars(r)["id"]
		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		err = uh.userRepository.DeleteUser(int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)
	}).Methods(http.MethodDelete)

	router.HandleFunc("/users/edit", func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		err = uh.userRepository.EditUser(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	}).Methods(http.MethodPost)

}
