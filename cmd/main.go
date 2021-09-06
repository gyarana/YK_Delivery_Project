package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nix_education/conf"
	"nix_education/pkg/handlers"
	"nix_education/pkg/model/DBrepositories"
)

const httpPort = ":8080"

func main() {
	db, err := conf.GetDB()
	if err != nil {
		log.Printf("Please check database connection " + err.Error())
		return
	}
	defer db.Close()
	router := mux.NewRouter()
	userDBRepository := DBrepositories.NewUserDBRepository(db)
	userHandler := handlers.NewUserHandler(userDBRepository)
	userHandler.InitHandleFuncRoutes(router)

	orderDBRepository := DBrepositories.NewOrderDBRepository(db)
	orderHandler := handlers.NewOrderHandler(orderDBRepository)
	orderHandler.InitOrdersHandleFuncRoutes(router)

	log.Fatal(http.ListenAndServe(httpPort, nil))

}
