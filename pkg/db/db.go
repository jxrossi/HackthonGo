package db

import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	StorageDB *sql.DB
)

func init() {
	dataSource := "root:root@tcp(localhost:3306)/fantasy_products"
	var err error
	StorageDB, err = sql.Open("mysql", dataSource); if err != nil {
		panic(err)
	}

	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("DB configured successfully")
}