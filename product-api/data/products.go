package data

import (
	"fmt"
	"time"
)

// Product defines the structure of an API product
type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Price        float32 `json:"price" validate:"required,gt=0"`
	SKU          string  `json:"sku" validate:"required,sku"`
	CreatedOn    string  `json:"_"`
	UpdatedOn    string  `json:"_"`
	DeletedOn    string  `json:"_"`
}

// ErrorProductNotFound is an error raised when a product is not found in the store
var ErrorProductNotFound = fmt.Errorf("Product Not Found")

// Products is a collection of Product
type Products []*Product

// GetProducts returns a list of the products
func GetProducts() Products {
	return productList
}

// GetProductByID returns a single product which matches the id from the store
// If a product is not found this function returns a ProductNotFound error
func GetProductByID(id int) (*Product, error) {
	index := findIndexByProductID(id)

	if index == -1 {
		return nil, ErrorProductNotFound
	}

	return productList[index], nil
}

// AddProduct adds a new product to the store
func AddProduct(product Product) {
	maxID := productList[len(productList)-1].ID
	product.ID = maxID + 1
	productList = append(productList, &product)
}


// UpdateProduct replaces a product in the store with the given item
// If a product with the given id does not exist in the database
// this function returns a ProductNotFound error
func UpdateProduct(product Product)  error {
	index := findIndexByProductID(product.ID)

	if index == -1 {
		return ErrorProductNotFound
	}

	productList[index] = &product

	return nil
}

// DeleteProduct removes product from the store
func DeleteProduct(id int) error{
	index := findIndexByProductID(id)

	if index == -1 {
		return ErrorProductNotFound
	}

	productList = append(productList[:index], productList[index+1])

	return nil
}

func getNextID() int {
	lastProduct := productList[len(productList) - 1]
	return lastProduct.ID + 1
}

// findIndex finds the index of a product in the store
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for index, product := range productList {
		if product.ID == id {
			return index
		}
	}
	return -1
}


// ProductList is a list of static product data
var productList = []*Product{
	&Product{
		ID:             1,
		Name:           "Latte",
		Description:    "Frothy, miky coffee",
		Price:          2.45,
		SKU:            "abc123",
		CreatedOn:      time.Now().UTC().String(),
		UpdatedOn:      time.Now().UTC().String(),
	},
	&Product{
		ID:               2,
		Name:            "Espresso",
		Description:     "Short and strong coffee without milk",
		Price:           1.99,
		SKU:             "zyx987",
		CreatedOn:       time.Now().UTC().String(),
		UpdatedOn:       time.Now().UTC().String(),
	},
}
