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
	QUERY_EXISTS_INVOICES = "SELECT EXISTS (SELECT 1 FROM invoices);"
	QUERY_EXISTS_SALES = "SELECT EXISTS (SELECT 1 FROM sales);"
	QUERY_STORE_CUSTOMERS = "INSERT INTO customers(id, last_name, first_name, condicion) VALUES(?, ?, ?, ?);"
	QUERY_STORE_PRODUCTS = "INSERT INTO products(id, description, price) VALUES(?, ?, ?);"
	QUERY_STORE_INVOICES = "INSERT INTO invoices(id, datetime, id_customer) VALUES(?, ?, ?);"
	QUERY_STORE_SALES = "INSERT INTO sales(id, id_invoice, id_product, quantity) VALUES(?, ?, ?, ?);"
)

type Repository interface {
	Store(table string, data interface{}) error
	ExistsCustomers() bool
	ExistsProducts() bool
	ExistsInvoices() bool
	ExistsSales() bool
}

type repository struct {}

func NewRepository(db *sql.DB) Repository {
	return &repository{}
}

func (repo *repository) Store(table string, data interface{}) error {
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

		case "invoices":
			invoicesStored := data.(models.Invoices)
			stmt, err := db.Prepare(QUERY_STORE_INVOICES)
			if err != nil {
				return err
			}

			defer stmt.Close()

			_, err = stmt.Exec(invoicesStored.ID, invoicesStored.DateTime, invoicesStored.IDCustomer)
			if err != nil {
				return err
			}
			return nil

		case "sales":
			salesStored := data.(models.Sales)
			stmt, err := db.Prepare(QUERY_STORE_SALES)
			if err != nil {
				return err
			}

			defer stmt.Close()

			_, err = stmt.Exec(salesStored.ID, salesStored.IDInvoice, salesStored.IDProduct, salesStored.Quantity)
			if err != nil {
				return err
			}
			return nil


		default:
			return errors.New("wrong table name")
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

func (repo *repository) ExistsInvoices() bool {
	db := db.StorageDB

	var cond bool
	err := db.QueryRow(QUERY_EXISTS_INVOICES).Scan(&cond); if err != nil {
		return false
	}
	return cond
}

func (repo *repository) ExistsSales() bool {
	db := db.StorageDB

	var cond bool
	err := db.QueryRow(QUERY_EXISTS_SALES).Scan(&cond); if err != nil {
		return false
	}
	return cond
}