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

type ProductHandlerI interface {
	GetMenuById(w http.ResponseWriter, r *http.Request)
	GetAllMenu(w http.ResponseWriter, r *http.Request)
	GetAllMenuByRestID(w http.ResponseWriter, r *http.Request)
}

type MenuHandler struct {
	menuService services.MenuServiceI
	logger      *logrus.Logger
}

func (m MenuHandler) GetMenuById(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":

		var menu model.Product
		err := json.NewDecoder(r.Body).Decode(&menu)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}
		menuS, err := m.menuService.GetAllMenuByRestID(menu.IDSupplier)

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

func (m MenuHandler) GetAllMenuByRestID(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		//reqProduct := new(model.ProductSupplierIDRequest)
		var menu model.Product
		err := json.NewDecoder(r.Body).Decode(&menu.IDSupplier)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotAcceptable)
			return
		}

		menuS, err := m.menuService.GetAllMenuByRestID(menu.IDSupplier)

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
