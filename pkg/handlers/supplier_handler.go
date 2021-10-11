package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"nix_education/model"
	"nix_education/services"
)

func NewSupplierHandler(supplierService services.SupplierServiceI, logger *logrus.Logger) *SupplierHandler {
	return &SupplierHandler{
		supplierService: supplierService,
		logger:          logger,
	}
}

type SupplierHandlerI interface {
	GetSupplierByID(w http.ResponseWriter, r *http.Request)
	GetAllSuppliers(w http.ResponseWriter, r *http.Request)
	GetSuppliersByType(w http.ResponseWriter, r *http.Request)
}

type SupplierHandler struct {
	supplierService services.SupplierServiceI
	logger          *logrus.Logger
}

func (s SupplierHandler) GetSupplierByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		var supplier model.Restaurant
		err := json.NewDecoder(r.Body).Decode(&supplier.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		supplier_serv, err := s.supplierService.GetByID(supplier.Id)

		if supplier.Id == 0 {
			http.Error(w, "no such supplier", http.StatusNotAcceptable)
			return
		}

		jSupplier, err := json.Marshal(*supplier_serv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jSupplier)
		if err != nil {
			s.logger.Fatal(err)
		}

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (s SupplierHandler) GetAllSuppliers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		suppliers, err := s.supplierService.GetAll()

		jSuppliers, err := json.Marshal(*suppliers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		length, err := w.Write(jSuppliers)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(length)

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}
func (s SupplierHandler) GetSuppliersByType(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		var supplier model.Restaurant
		err := json.NewDecoder(r.Body).Decode(&supplier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		supplier_serv, err := s.supplierService.GetAllByType(supplier.Type)

		jSuppliers, err := json.Marshal(*supplier_serv)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jSuppliers)
		if err != nil {
			s.logger.Fatal(err)
		}

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}
