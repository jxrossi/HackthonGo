package main

import (
	internal "HacktonGo/internal"
	data "HacktonGo/pkg/data_handler"
	db "HacktonGo/pkg/db"
	"encoding/json"
	"log"
    "fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
    db := db.StorageDB

    repo := internal.NewRepository(db)
	service := internal.NewService(repo)
    
    //insert data into tables
    log.Println(data.InsertData("customers", service))
    log.Println(data.InsertData("products", service))
    log.Println(data.InsertData("invoices", service))
    log.Println(data.InsertData("sales", service))

    //insert totals calculated into invoices table
    log.Println(data.InsertTotalInvoices(service))

    fmt.Println("------------")

    //first exercise
    log.Println("First Exercise")
    dataBytes1, _ := service.Enunciado1()
    dataJson1, _ := json.MarshalIndent(dataBytes1, "", "   ")
    fmt.Println(string(dataJson1))

    //second exercise
    log.Println("Second Exercise")
    dataBytes2, _ := service.Enunciado2()
    dataJson2, _ := json.MarshalIndent(dataBytes2, "", "   ")
    fmt.Println(string(dataJson2))

    //third exercise
    log.Println("Third Exercise")
    dataBytes3, _ := service.Enunciado3()
    dataJson3, _ := json.MarshalIndent(dataBytes3, "", "   ")
    fmt.Println(string(dataJson3))
}