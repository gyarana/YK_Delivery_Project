package conf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbName := "delivery_db"
	dbUser := "root"
	dbPass := "yaros08"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return

}
