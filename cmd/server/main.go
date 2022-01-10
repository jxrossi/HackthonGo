package main

import (
	"log"
	_ "github.com/go-sql-driver/mysql"
	internal "HacktonGo/internal"
	db "HacktonGo/pkg/db"
    data "HacktonGo/pkg/data_handler"
)

func main() {
    db := db.StorageDB

    repo := internal.NewRepository(db)
	service := internal.NewService(repo)
    
    log.Println(data.InsertData("customers", service))
    log.Println(data.InsertData("products", service))
    log.Println(data.InsertData("invoices", service))
    log.Println(data.InsertData("sales", service))
}