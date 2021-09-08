package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"nix_education/model"
	"nix_education/model/repositories"
	"strconv"
)

type OrderHandler struct {
	orderDBRepository repositories.OrderRepositoryI
}

func NewOrderHandler(orderRepo repositories.OrderRepositoryI) *OrderHandler {
	return &OrderHandler{orderDBRepository: orderRepo}
}

func (oh OrderHandler) InitOrdersHandleFuncRoutes(router *mux.Router) {

	router.HandleFunc("/orders/create", func(w http.ResponseWriter, r *http.Request) {
		var order model.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		err = oh.orderDBRepository.CreateOrder(&order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)

	}).Methods(http.MethodPost)

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		orders, err := oh.orderDBRepository.GetAllOrders()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(orders)
	}).Methods(http.MethodGet)

	router.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		strId := mux.Vars(r)["id"]
		id, err := strconv.Atoi(strId)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		order, err := oh.orderDBRepository.GetOrder(int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}).Methods(http.MethodGet)

	router.HandleFunc("/orders/{id}", func(w http.ResponseWriter, r *http.Request) {
		strId := mux.Vars(r)["id"]
		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		err = oh.orderDBRepository.DeleteOrder(int32(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(true)
	}).Methods(http.MethodDelete)

	router.HandleFunc("/orders/edit", func(w http.ResponseWriter, r *http.Request) {
		var order model.Order
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		err = oh.orderDBRepository.EditOrder(&order)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(order)
	}).Methods(http.MethodPost)

}
