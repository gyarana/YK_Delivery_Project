package handlers

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"nix_education/model"
	"nix_education/services"
)

func NewMenuHandler(menuService services.MenuServiceI, logger *logrus.Logger) *MenuHandler {
	return &MenuHandler{
		menuService: menuService,
		logger:      logger,
	}
}

type MenuHandlerI interface {
	GetAllMenuByID(w http.ResponseWriter, r *http.Request)
	GetAllMenu(w http.ResponseWriter, r *http.Request)
	GetAllMenuByRestID(w http.ResponseWriter, r *http.Request)
	CreateMenu(w http.ResponseWriter, r *http.Request)
	UpdateMenu(w http.ResponseWriter, r *http.Request)
	DeleteMenu(w http.ResponseWriter, r *http.Request)
}

type MenuHandler struct {
	menuService services.MenuServiceI
	logger      *logrus.Logger
}

func (m MenuHandler) GetAllMenuByRestID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":

		var supID model.MenuSupplierIDRequest
		err := json.NewDecoder(r.Body).Decode(&supID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		menuS, err := m.menuService.GetAllMenuByRestID(supID.SupplierID)

		if menuS == nil {
			http.Error(w, "Non-existent position", http.StatusNotAcceptable)
			return
		}

		jMenu, err := json.Marshal(*menuS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jMenu)
		if err != nil {
			m.logger.Fatal(err)
		}

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (m MenuHandler) GetAllMenu(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		menu, err := m.menuService.GetAllMenu()

		jMenu, err := json.Marshal(*menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jMenu)
		if err != nil {
			m.logger.Fatal(err)
		}

	default:
		http.Error(w, "Only GET is Allowed", http.StatusMethodNotAllowed)
	}
}

func (m MenuHandler) GetAllMenuByID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var menuID model.MenuRequest
		err := json.NewDecoder(r.Body).Decode(&menuID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		menuS, err := m.menuService.GetAllMenuByRestID(menuID.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		if len(*menuS) == 0 {
			http.Error(w, "Non exist product for this supplier. Please check ID", http.StatusNotAcceptable)
			return
		}

		jMenu, err := json.Marshal(*menuS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write(jMenu)
		if err != nil {
			m.logger.Fatal(err)
		}

	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}
}

func (m MenuHandler) CreateMenu(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var menu model.Product
		err := json.NewDecoder(r.Body).Decode(&menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = m.menuService.CreateMenu(&menu)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

}

func (m MenuHandler) UpdateMenu(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "PUT":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var menu model.Product
		err := json.NewDecoder(r.Body).Decode(&menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = m.menuService.UpdateMenu(&menu)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

}

func (m MenuHandler) DeleteMenu(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		var menuID model.MenuRequest
		err := json.NewDecoder(r.Body).Decode(&menuID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		err = m.menuService.DeleteMenu(menuID.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		w.WriteHeader(http.StatusOK)
	default:
		http.Error(w, "Only POST is Allowed", http.StatusMethodNotAllowed)
	}

}
