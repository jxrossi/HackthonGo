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
	QUERY_EXISTS_TOTAL_INVOICES = "SELECT EXISTS (SELECT total from invoices WHERE total > '');"


	QUERY_STORE_CUSTOMERS = "INSERT INTO customers(id, last_name, first_name, condicion) VALUES(?, ?, ?, ?);"
	QUERY_STORE_PRODUCTS = "INSERT INTO products(id, description, price) VALUES(?, ?, ?);"
	QUERY_STORE_INVOICES = "INSERT INTO invoices(id, datetime, id_customer, total) VALUES(?, ?, ?, ?);"
	QUERY_STORE_SALES = "INSERT INTO sales(id, id_invoice, id_product, quantity) VALUES(?, ?, ?, ?);"


	QUERY_GET_TOTAL_INVOICES = `SELECT SUM(s.quantity * p.price) AS total FROM invoices i
	INNER JOIN sales s ON i.id = s.id_invoice
	INNER JOIN products p ON s.id_product = p.id
	WHERE i.id = ?
	GROUP BY i.id;`


	QUERY_STORE_TOTAL_INVOICES = "UPDATE invoices SET total = ? WHERE id = ?;"


	QUERY_ENUNCIADO_1 = `SELECT c.condicion, ROUND(SUM(s.quantity * p.price), 2) AS total FROM invoices i
						 INNER JOIN sales s ON i.id = s.id_invoice
						 INNER JOIN products p ON s.id_product = p.id
						 INNER JOIN customers c ON i.id_customer = c.id
						 GROUP BY c.condicion`

	QUERY_ENUNCIADO_2 = `SELECT p.description, ROUND(SUM(s.quantity), 2) AS total FROM sales s
						 INNER JOIN products p ON s.id_product = p.id
						 GROUP BY p.description
						 ORDER BY total DESC
						 LIMIT 5;`

	QUERY_ENUNCIADO_3 = `SELECT c.last_name, c.first_name FROM customers c
						 INNER JOIN invoices i ON c.id = i.id_customer
						 INNER JOIN sales s ON i.id = s.id_invoice
						 INNER JOIN products p ON s.id_product = p.id
						 ORDER BY p.price, c.last_name
						 LIMIT 5;`

)

type Repository interface {
	Store(table string, data interface{}) error
	ExistsCustomers() bool
	ExistsProducts() bool
	ExistsInvoices() bool
	ExistsSales() bool
	ExistsTotalsInvoices() bool
	StoreTotalInvoices(id int) error
	Enunciado1() ([]models.Enunciado1, error)
	Enunciado2() ([]models.Enunciado2, error)
	Enunciado3() ([]models.Enunciado3, error)
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

func (repo *repository) ExistsTotalsInvoices() bool {
	db := db.StorageDB

	var cond bool
	err := db.QueryRow(QUERY_EXISTS_TOTAL_INVOICES).Scan(&cond); if err != nil {
		return false
	}
	return cond
}

func (repo *repository) StoreTotalInvoices(id int) error {
	db := db.StorageDB
	var aaa float64

	err := db.QueryRow(QUERY_GET_TOTAL_INVOICES, id).Scan(&aaa)
	if err != nil || err == sql.ErrNoRows {
		return err
	}

	stmt, err := db.Prepare(QUERY_STORE_TOTAL_INVOICES); if err != nil {
		return err
	}
	
	_, err = stmt.Exec(aaa, id); if err != nil {
		return err
	}
	return nil
}

func (repo *repository) Enunciado1() ([]models.Enunciado1, error) {
	db := db.StorageDB
	
	var enun models.Enunciado1
	var enuns []models.Enunciado1 

	rows, err := db.Query(QUERY_ENUNCIADO_1)
    if err != nil {
		return nil, err
    }
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		if err := rows.Scan(&enun.Condition, &enun.Total); err != nil {
			return nil, err
		}
		enuns = append(enuns, enun)
	}

	return enuns, nil
}

func (repo *repository) Enunciado2() ([]models.Enunciado2, error) {
	db := db.StorageDB
	
	var enun models.Enunciado2
	var enuns []models.Enunciado2

	rows, err := db.Query(QUERY_ENUNCIADO_2)
    if err != nil {
		return nil, err
    }
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		if err := rows.Scan(&enun.Description, &enun.Total); err != nil {
			return nil, err
		}
		enuns = append(enuns, enun)
	}

	return enuns, nil
}

func (repo *repository) Enunciado3() ([]models.Enunciado3, error) {
	db := db.StorageDB
	
	var enun models.Enunciado3
	var enuns []models.Enunciado3

	rows, err := db.Query(QUERY_ENUNCIADO_3)
    if err != nil {
		return nil, err
    }
	if err != nil {
		return nil, err
	}
	
	for rows.Next() {
		if err := rows.Scan(&enun.LastName, &enun.FirstName); err != nil {
			return nil, err
		}
		enuns = append(enuns, enun)
	}

	return enuns, nil
}