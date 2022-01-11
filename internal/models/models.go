package models

type Customers struct {
	ID        int    `json:"id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
	Condition string `json:"condition"`
}

type Products struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Invoices struct {
	ID         int     `json:"id"`
	DateTime   string  `json:"datetime"`
	IDCustomer int     `json:"id_customer"`
	Total      float64 `json:"total"`
}

type Sales struct {
	ID        int     `json:"id"`
	IDInvoice int     `json:"id_invoice"`
	IDProduct int     `json:"id_product"`
	Quantity  float64 `json:"quantity"`
}

type Enunciado1 struct {
	Condition string  `json:"condition"`
	Total     float64 `json:"total"`
}

type Enunciado2 struct {
	Description string  `json:"description"`
	Total       float64 `json:"total"`
}

type Enunciado3 struct {
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}
