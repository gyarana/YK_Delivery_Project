package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"nix_education/model"
	"nix_education/services"
)

func NewOrderHandler(orderService services.OrderServiceI, logger *logrus.Logger) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
		logger:       logger,
	}
}

type OrderHandlerI interface {
	GetAllOrder(w http.ResponseWriter, r *http.Request)
	GetOrder(w http.ResponseWriter, r *http.Request)
	CreateOrder(w http.ResponseWriter, r *http.Request)
	UpdateOrder(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
}

type OrderHandler struct {
	orderService services.OrderServiceI
	logger       *logrus.Logger
}

func (o OrderHandler) GetAllOrder(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		orders, err := o.orderService.GetAllOrders()

		jOrder, err := json.Marshal(*orders)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jOrder)
		if err != nil {
			o.logger.Fatal(err)
		}

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}

}

func (o OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var orderID model.OrderRequestID
		err := json.NewDecoder(r.Body).Decode(&orderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		orderS, err := o.orderService.GetOrder(orderID.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		jOrder, err := json.Marshal(*orderS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jOrder)
		if err != nil {
			o.logger.Fatal(err)
		}

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (o OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var order model.OrderRequest
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = o.orderService.CreateOrder(&order)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

}

func (o OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var order model.OrderRequest
		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = o.orderService.EditOrder(&order)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

}

func (o OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var orderID model.OrderRequestID
		err := json.NewDecoder(r.Body).Decode(&orderID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = o.orderService.DeleteOrder(orderID.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

}
