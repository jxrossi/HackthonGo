package db

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	StorageDB *sql.DB
)

func init() {
	_ = godotenv.Load()

	dataSource := os.ExpandEnv("${U}:${PW}@tcp(localhost:3306)/${DB_NAME}")
	
	var err error
	StorageDB, err = sql.Open("mysql", dataSource); if err != nil {
		panic(err)
	}

	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

	log.Println("DB configured successfully")
}