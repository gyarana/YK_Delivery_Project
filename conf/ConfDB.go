package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := os.Getenv("dbDriver")
	dbName :=  os.Getenv("dbName")
	dbUser :=  os.Getenv("dbUser")
	dbPass :=  os.Getenv("dbPass")
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return

}
