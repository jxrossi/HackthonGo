package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	data "HacktonGo/internal"
	"HacktonGo/internal/models"
	"HacktonGo/pkg/db"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
    var (
        file *os.File
        part []byte
        prefix bool
    )
    if file, err = os.Open(path); err != nil {
        return
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    buffer := bytes.NewBuffer(make([]byte, 0))
    for {
        if part, prefix, err = reader.ReadLine(); err != nil {
            break
        }
        buffer.Write(part)
        if !prefix {
            lines = append(lines, buffer.String())
            buffer.Reset()
        }
    }
    if err == io.EOF {
        err = nil
    }
    return
}

//Only inserts Costumers Data if table is empty.
func insertData(table string, service data.Service) {
        
        //Read Data
        lines, err := readLines("../../datos/" + table + ".txt"); if err != nil {
            fmt.Printf("Error: %s\n", err)
        }

        switch table {
            case "customers":
                //Check if Customers Table already has data.
                if !(service.ExistsCustomers()) {
                    var custom models.Customers
                    for _, line := range lines {
                        v := strings.Split(line, "#$%#")
            
                        id, _ := strconv.Atoi(v[0])
                        custom.ID = id
                        custom.LastName = v[1]
                        custom.FirstName = v[2]
                        custom.Condition = v[len(v)-1]
            
                        err := service.StoreCustomers(table, custom); if err != nil {
                            fmt.Printf("Error: %s", err)
                        }
                    }
                }

            case "products":
                if !(service.ExistsProducts()) {
                    for _, line := range lines {
                        v := strings.Split(line, "#$%#")
            
                        var prods models.Products
            
                        id, _ := strconv.ParseInt(v[0], 10, 64)
                        prods.ID = id
                        prods.Description = v[1]
                        price, _ := strconv.ParseFloat(v[len(v)-1], 64)
                        prods.Price = price;
                        
                        err := service.StoreCustomers(table, prods); if err != nil {
                            fmt.Printf("Error: %s", err)
                        }
                    }
                }
                
            default:
                return 
        }
}

//Only inserts Products Data if table is empty.
/* func insertProductsData(service data.Service) {
    //Check if Customers Table already has data.
    if !(service.ExistsProducts()) {
        //Read Customers Data
        lines, err := readLines("../../datos/products.txt")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }

        for _, line := range lines {
            v := strings.Split(line, "#$%#")

            var custom customer.Customers

            id, _ := strconv.ParseInt(v[0], 10, 64)
            custom.ID = id
            custom.LastName = v[1]
            custom.FirstName = v[2]
            custom.Condition = v[len(v)-1]

            err := service.StoreCustomers("customers", custom)
            if err != nil {
                fmt.Printf("Error: %s", err)
            }
            
        }
    }
} */

func main() {

    db := db.StorageDB

    repo := data.NewRepository(db)
	service := data.NewService(repo)
    
    insertData("customers", service)
    log.Println("Executed insertData Products successfully")

    /* table := "products"
    lines, err := readLines("../../datos/" + table + ".txt")
        if err != nil {
            fmt.Printf("Error: %s\n", err)
            return
        }

        for _, line := range lines {
            v := strings.Split(line, "#$%#")

            var prods customer.Products

            id, _ := strconv.ParseInt(v[0], 10, 64)
            prods.ID = id
            prods.Description = v[1]
            price, _ := strconv.ParseFloat(v[len(v)-1], 64)
            prods.Price = price;
            
            fmt.Println(prods.Price);
        } */


    insertData("products", service)
    log.Println("Executed insertData Products successfully")
    
    //Leer archivo Products
    /* lines, err = readLines("../../datos/products.txt")
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    } */

    //TODO: Insertar linea x linea Products
    /* for _, line := range lines {
        fmt.Println(line)
    } */

    //Leer archivo Invoices
    /* lines, err = readLines("../../datos/invoices.txt")
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    } */

    //TODO: Insertar linea x linea Invoices
    /* for _, line := range lines {
        fmt.Println(line)
    } */

    //Leer archivo Sales
    /* lines, err = readLines("../../datos/sales.txt")
    if err != nil {
        fmt.Printf("Error: %s\n", err)
        return
    } */

    //TODO: Insertar linea x linea Sales
    /* for _, line := range lines {
        fmt.Println(line)
    } */
}