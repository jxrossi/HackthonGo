package models

type Customers struct {
	ID       		int    			`json:"id"`
	LastName 		string 			`json:"last_name"`
	FirstName   	string 			`json:"first_name"`
	Condition     	string    		`json:"condition"`
}

type Products struct {
	ID       			int64    		`json:"id"`
	Description 		string 			`json:"description"`
	Price   			float64 		`json:"price"`
}

type Invoices struct {
	ID       			int64    		`json:"id"`
	DateTime			string 			`json:"datetime"`
	IDCustomer    		int64 			`json:"id_customer"`
	Total 				float64 		`json:"total"`
}

type Sales struct {
	ID 				int64 			`json:"id"`
	IDInvoice 		int64 			`json:"id_invoice"`
	IDProduct 		int64 			`json:"id_product"`
	Quantity 		float64 		`json:"quantity"`
}