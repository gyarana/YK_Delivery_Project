package main

import (
	_ "github.com/go-sql-driver/mysql"
	"nix_education/cmd/app"
)

const (
	urlRest  = "http://foodapi.true-tech.php.nixdev.co/suppliers"
	urlItems = "http://foodapi.true-tech.php.nixdev.co/suppliers/%v/menu"
)

func main() {
	app.Start()
}
