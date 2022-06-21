package config

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)

func GetDB() (db *sql.DB, err error) {
	db, err1 := sql.Open("sqlserver", "server=LAPTOP-NQDHTU17\\SQLEXPRESS; database=Battleground;")
	return db, err1
}
