package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"nix_education/model"
	"nix_education/services"
)

func NewCartHandler(cartService services.CartServiceI, logger *logrus.Logger) *CartHandler {
	return &CartHandler{
		cartService: cartService,
		logger:      logger,
	}
}

type CartHandlerI interface {
	GetCartByID(w http.ResponseWriter, r *http.Request)
	CreateCart(w http.ResponseWriter, r *http.Request)
	UpdateCart(w http.ResponseWriter, r *http.Request)
	DeleteCart(w http.ResponseWriter, r *http.Request)
}

type CartHandler struct {
	cartService services.CartServiceI
	logger      *logrus.Logger
}

func (c CartHandler) GetCartByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var cartID model.CartRequest
		err := json.NewDecoder(r.Body).Decode(&cartID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		_, err = c.cartService.GetCartByID(cartID.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (c CartHandler) CreateCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var cart model.Cart
		err := json.NewDecoder(r.Body).Decode(&cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = c.cartService.CreateCart(&cart)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (c CartHandler) UpdateCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "PUT":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var cart model.Cart
		err := json.NewDecoder(r.Body).Decode(&cart)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = c.cartService.UpdateCart(&cart)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only PUT is Allowed", http.StatusMethodNotAllowed)
	}

}

func (c CartHandler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var cartID model.CartRequest
		err := json.NewDecoder(r.Body).Decode(&cartID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = c.cartService.DeleteCart(cartID.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}
