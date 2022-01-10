package data_handler

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"HacktonGo/internal/models"
	internal "HacktonGo/internal"
)

// Read a whole file into the memory and store it as array of lines
func ReadLines(path string) (lines []string, err error) {
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

//Only inserts data if table is empty.
func InsertData(table string, service internal.Service) string {
	//Read Data
	lines, err := ReadLines("../../datos/" + table + ".txt"); if err != nil {
		return fmt.Sprintf("Error: %s\n", err)
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
		
					err := service.Store(table, custom); if err != nil {
						return fmt.Sprintf("Error: %s\n", err)
					}
				}
				return "Data Customers inserted successfully"
			}
			return "There is already data in Customers table"

		case "products":
			if !(service.ExistsProducts()) {
				for _, line := range lines {
					v := strings.Split(line, "#$%#")
		
					var prods models.Products
		
					id, _ := strconv.Atoi(v[0])
					prods.ID = id
					prods.Description = v[1]
					price, _ := strconv.ParseFloat(v[len(v)-1], 64)
					prods.Price = price;
					
					err := service.Store(table, prods); if err != nil {
						return fmt.Sprintf("Error: %s\n", err)
					}
				}
				return "Data Products inserted successfully"
			}
			return "There is already data in Products table"

		case "invoices":
			var invs models.Invoices
			if !(service.ExistsInvoices()) {
				for _, line := range lines {
					v := strings.Split(line, "#$%#")
		
					id, _ := strconv.Atoi(v[0])
					invs.ID = id
					invs.DateTime = v[1]
					idCustomer, _ := strconv.Atoi(v[2])
					invs.IDCustomer = idCustomer

					err := service.Store(table, invs); if err != nil {
						return fmt.Sprintf("Error: %s\n", err)
					}
				}
				return "Data Invoices inserted successfully"
			}
			return "There is already data in Invoices table"

		case "sales":
			if !(service.ExistsSales()) {
				for _, line := range lines {
					v := strings.Split(line, "#$%#")
		
					var sales models.Sales
		
					id, _ := strconv.Atoi(v[0])
					sales.ID = id

					idInvoice, _ := strconv.Atoi(v[1])
					sales.IDInvoice = idInvoice

					idProduct, _ := strconv.Atoi(v[2])
					sales.IDProduct = idProduct

					quantity, _ := strconv.ParseFloat(v[len(v)-1], 64)
					sales.Quantity = quantity

					err := service.Store(table, sales); if err != nil {
						return fmt.Sprintf("Error: %s\n", err)
					}
				}
				return "Data Sales inserted successfully"
			}
			return "There is already data in Sales table"
	}
	return "Wrong given table name"
}

func InsertTotalInvoices(service internal.Service) string {
	//Read Data
	lines, err := ReadLines("../../datos/invoices.txt"); if err != nil {
		return fmt.Sprintf("Error: %s\n", err)
	}

	if !(service.ExistsTotalsInvoices()) {
		if service.ExistsSales() {
			for _, line := range lines {
				v := strings.Split(line, "#$%#")
				id, _ := strconv.Atoi(v[0])
	
				err := service.StoreTotalInvoices(id); if err != nil {
					return fmt.Sprintf("Error: %s\n", err)
				}
			}
			return "Stored Invoice totals fields successfully"
		}
		return "Sales table is empty or there is , try populating it and try again"
	}
	return "Totals from Invoices already calculated"
	
}

