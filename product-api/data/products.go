package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure of an API product
type Product struct {
	ID					int			`json:"id"`
	Name				string	`json:"name"`
	Description	string	`json:"description"`
	Price				float32	`json:"price"`
	SKU					string	`json:"sku"`
	CreatedOn		string	`json:"_"`
	UpdatedOn		string	`json:"_"`
	DeletedOn		string	`json:"_"`
}

// Products is a collection of Product
type Products []*Product

// ToJSON serializes contents of the collection to JSON using json package's NeWEncoder
// https://golang.org/pkg/encoding/json/#NewEncoder
func (p*Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts returns a list of the products
func GetProducts () Products {
	return ProductList
}

// ProductList is a list of static product data
var ProductList = []*Product{
	&Product{
		ID:						1,
		Name: 				"Latte",
		Description:	"Frothy, miky coffee",
		Price:				2.45,
		SKU: 					"abc123",
		CreatedOn:		time.Now().UTC().String(),
		UpdatedOn:		time.Now().UTC().String(),
	},
	&Product{
		ID:						2,
		Name: 				"Espresso",
		Description:	"Short and strong coffee without milk",
		Price:				1.99,
		SKU: 					"zyx987",
		CreatedOn:		time.Now().UTC().String(),
		UpdatedOn:		time.Now().UTC().String(),
	},
}
