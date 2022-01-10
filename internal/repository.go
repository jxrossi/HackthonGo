package internal

import (
	"HacktonGo/internal/models"
	"HacktonGo/pkg/db"
	"database/sql"
	"errors"
)

const (
	QUERY_EXISTS_CUSTOMERS = "SELECT EXISTS (SELECT 1 FROM customers);"
	QUERY_EXISTS_PRODUCTS = "SELECT EXISTS (SELECT 1 FROM products);"
	QUERY_STORE_CUSTOMERS = "INSERT INTO customers(id, last_name, first_name, condicion) VALUES(?, ?, ?, ?);"
	QUERY_STORE_PRODUCTS = "INSERT INTO products(id, description, price) VALUES(?, ?, ?);"
)


type Repository interface {
	StoreCustomers(table string, data interface{}) error
	ExistsCustomers() bool
	ExistsProducts() bool
}

type repository struct {}

func NewRepository(db *sql.DB) Repository {
	return &repository{}
}

func (repo *repository) StoreCustomers(table string, data interface{}) error {
	db := db.StorageDB

	switch table {
		case "customers":
			customerStored := data.(models.Customers)
			stmt, err := db.Prepare(QUERY_STORE_CUSTOMERS)
			if err != nil {
				return err
			}

			defer stmt.Close()

			_, err = stmt.Exec(customerStored.ID, customerStored.LastName, customerStored.FirstName, customerStored.Condition)
			if err != nil {
				return err
			}
			return nil

		case "products":
			productsStored := data.(models.Products)
			stmt, err := db.Prepare(QUERY_STORE_PRODUCTS)
			if err != nil {
				return err
			}

			defer stmt.Close()

			_, err = stmt.Exec(productsStored.ID, productsStored.Description, productsStored.Price)
			if err != nil {
				return err
			}
			return nil

		default:
			return errors.New("asdasd")
	}
}

func (repo *repository) ExistsCustomers() bool {
	db := db.StorageDB

	var cond bool
	err := db.QueryRow(QUERY_EXISTS_CUSTOMERS).Scan(&cond); if err != nil {
		return false
	}
	return cond
}

func (repo *repository) ExistsProducts() bool {
	db := db.StorageDB

	var cond bool
	err := db.QueryRow(QUERY_EXISTS_PRODUCTS).Scan(&cond); if err != nil {
		return false
	}
	return cond
}

/* personaLeida := Empleado{}

	err = db.QueryRow("SELECT * FROM personas LIMIT 1").Scan(&personaLeida.ID, &personaLeida.Nombre, &personaLeida.Apellido, &personaLeida.Edad, &personaLeida.IDCiudad)
	if err != nil || err == sql.ErrNoRows {
		fmt.Println(Empleado{})
	}

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    fmt.Println(cons) */