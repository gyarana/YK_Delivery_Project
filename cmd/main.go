package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"nix_education/conf"
	"nix_education/model/repositories"
	"nix_education/parser"
)

const (
	urlRest  = "http://foodapi.true-tech.php.nixdev.co/suppliers"
	urlItems = "http://foodapi.true-tech.php.nixdev.co/suppliers/%v/menu"
)

func main() {
	var logger *logrus.Logger
	db, err := conf.GetDB()
	if err != nil {
		logrus.Error(err)
	}
	suppliersRepository := repositories.NewRestaurantsRepository(db)
	menuRepository := repositories.NewMenuRepository(db, logger)
	menuParser := parser.NewRestarauntsParser(urlRest, urlItems, logger, suppliersRepository, menuRepository)
	menuParser.TimeFieldUpdate()
	db.Close()
}
