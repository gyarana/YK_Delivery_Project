package app

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"nix_education/conf"
	"nix_education/middleware"
	"nix_education/model/repositories"
	"nix_education/parser"
	"nix_education/pkg/handlers"
	"nix_education/server"
	"nix_education/services"
	"os"
	"os/signal"
	"syscall"
)

//url format
//const (
//	urlRest  = "http://foodapi.true-tech.php.nixdev.co/suppliers"
//	urlItems = "http://foodapi.true-tech.php.nixdev.co/suppliers/%v/menu"
//)

func Start() {
	var logger *logrus.Logger
	db, err := conf.GetDB()
	if err != nil {
		logger.Fatal(err.Error())
	}

	menuRepo := repositories.NewMenuRepository(db, logger)
	menuService := services.NewProductService(menuRepo)
	menuHandler := handlers.NewMenuHandler(menuService, logger)
	cartRepo := repositories.NewCartRepository(db, logger)
	cartService := services.NewCartService(cartRepo)
	cartHandler := handlers.NewCartHandler(cartService, logger)
	suppRepo := repositories.NewRestaurantsRepository(db, logger)
	supplierService := services.NewSupplierService(suppRepo)
	supplierHandler := handlers.NewSupplierHandler(supplierService, logger)
	orderRepo := repositories.NewOrderRepository(db)
	orderService := services.NewOrderService(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderService, logger)
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	tokenService := services.NewTokenService()
	userHandler := handlers.NewLoginHandler(userService, tokenService)
	authHandler := middleware.NewAuthMiddlware(tokenService)

	mux := http.NewServeMux()
	mux.HandleFunc("/getmenu", menuHandler.GetAllMenu)
	mux.HandleFunc("/createmenu", menuHandler.CreateMenu)
	mux.HandleFunc("/getmenubyid", menuHandler.GetAllMenuByID)
	mux.HandleFunc("/getmenubyrest", menuHandler.GetAllMenuByRestID)
	mux.HandleFunc("/deletemenu", menuHandler.DeleteMenu)
	mux.HandleFunc("/updatemenu", menuHandler.UpdateMenu)
	//
	mux.HandleFunc("/getcart", cartHandler.GetCartByID)
	mux.HandleFunc("/createcart", cartHandler.CreateCart)
	mux.HandleFunc("/deletecart", cartHandler.DeleteCart)
	mux.HandleFunc("/updatecart", cartHandler.UpdateCart)
	//
	mux.HandleFunc("/getallrest", supplierHandler.GetAllSuppliers)
	mux.HandleFunc("/getrestbyid", supplierHandler.GetSupplierByID)
	mux.HandleFunc("/getrestbytype", supplierHandler.GetSuppliersByType)
	mux.HandleFunc("/createrest", supplierHandler.CreateSupplier)
	mux.HandleFunc("/deleterest", supplierHandler.DeleteSupplier)
	mux.HandleFunc("/updaterest", supplierHandler.UpdateSupplier)
	//
	mux.HandleFunc("/getorderbyid", orderHandler.GetOrder)
	mux.HandleFunc("/getallorder", orderHandler.GetAllOrder)
	mux.HandleFunc("/createorder", orderHandler.CreateOrder)
	mux.HandleFunc("/deleteorder", orderHandler.DeleteOrder)
	mux.HandleFunc("/updateorder", orderHandler.UpdateOrder)
	//
	mux.HandleFunc("/getprofile", authHandler.AccessTokenCheck(userHandler.GetUserProfile))
	mux.HandleFunc("/login", authHandler.AccessTokenCheck(userHandler.Login))
	mux.HandleFunc("/createuser", authHandler.AccessTokenCheck(userHandler.CreateNewUser))
	mux.HandleFunc("/edituser", authHandler.AccessTokenCheck(userHandler.EditUserProfile))
	mux.HandleFunc("/refresh", authHandler.RefreshTokenCheck(userHandler.Refresh))

	menuParser := parser.NewRestarauntsParser(os.Getenv("urlRest"), os.Getenv("urlItems"), logger, suppRepo, menuRepo)
	go menuParser.TimeFieldUpdate()

	srv := new(server.Server)
	//
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
	//
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatalf("Server shutting down error: ", err.Error())
	}
	if err := db.Close(); err != nil {
		logger.Fatalf("Database connection closing error: ", err.Error())
	}
}
