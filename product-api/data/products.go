package data

import (
	"time"
)

// Product defines the structure of an API product
type Product struct {
	ID					int
	Name				string
	Description	string
	Price				float32
	SKU					string
	CreatedOn		string
	UpdatedOn		string
	DeletedOn		string
}

// GetProducts returns a list of the products
func GetProducts () []*Product {
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
