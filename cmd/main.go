package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"nix_education/conf"
	"nix_education/model/repositories"
	"nix_education/parser"
)

const (
	urlRest  = "http://foodapi.true-tech.php.nixdev.co/restaurants"
	urlItems = "http://foodapi.true-tech.php.nixdev.co/restaurants/%v/menu"
)

func main() {
	db, err := conf.GetDB()
	if err != nil {
		fmt.Println("panic")
	}
	suppliersRepository := repositories.NewRestaurantsRepository(db)
	menuRepository := repositories.NewMenuRepository(db)
	menuParser := parser.NewRestarauntsParser(urlRest, urlItems, suppliersRepository, menuRepository)
	menuParser.TimeFieldUpdate()
	db.Close()
}
