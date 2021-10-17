package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"nix_education/conf"
	"nix_education/model/repositories"
	"nix_education/pkg/handlers"
	"nix_education/server"
	"nix_education/services"
	"os"
	"os/signal"
	"syscall"
)

func Start() {
	var logger *logrus.Logger
	db, err := conf.GetDB()
	if err != nil {
		logger.Fatal(err.Error())
	}

	menuRepo := repositories.NewMenuRepository(db, logger)
	ms := services.NewProductService(menuRepo)
	mh := handlers.NewMenuHandler(ms, logger)
	cartRepo := repositories.NewCartRepository(db, logger)
	cs := services.NewCartService(cartRepo)
	ch := handlers.NewCartHandler(cs, logger)
	suppRepo := repositories.NewRestaurantsRepository(db, logger)
	ss := services.NewSupplierService(suppRepo)
	sh := handlers.NewSupplierHandler(ss, logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/getmenu", mh.GetAllMenu)
	mux.HandleFunc("/createmenu", mh.CreateMenu)
	mux.HandleFunc("/getmenubyid", mh.GetAllMenuByID)
	mux.HandleFunc("/getmenubyrest", mh.GetAllMenuByRestID)
	mux.HandleFunc("/deletemenu", mh.DeleteMenu)
	mux.HandleFunc("/updatemenu", mh.UpdateMenu)
	//
	//
	mux.HandleFunc("/getcart", ch.GetCartByID)
	mux.HandleFunc("/createcart", ch.CreateCart)
	mux.HandleFunc("/deletecart", ch.DeleteCart)
	mux.HandleFunc("/updatecart", ch.UpdateCart)
	//
	mux.HandleFunc("/getallrest", sh.GetAllSuppliers)
	mux.HandleFunc("/getrestbyid", sh.GetSupplierByID)
	mux.HandleFunc("/getrestbytype", sh.GetSuppliersByType)
	//
	mux.HandleFunc("/createrest", sh.CreateSupplier)
	mux.HandleFunc("/deleterest", sh.DeleteSupplier)
	mux.HandleFunc("/updaterest", sh.UpdateSupplier)

	srv := new(server.Server)

	go func() {
		if err := srv.StartServer(os.Getenv("port"), mux); err != nil {
			logger.Fatalf("Server running error: ", err.Error())
		}
	}()
	logger.Infof("Server started...")
	ctx := context.Background()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logger.Infof("App shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server shutting down error: ", err.Error())
	}
	if err := db.Close(); err != nil {
		logger.Fatalf("Database connection closing error: ", err.Error())
	}

}
