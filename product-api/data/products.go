package data

import (
	"fmt"
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure of an API product
type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	Price        float32 `json:"price"`
	SKU          string  `json:"sku"`
	CreatedOn    string  `json:"_"`
	UpdatedOn    string  `json:"_"`
	DeletedOn    string  `json:"_"`
}

// FromJSON decodes json serialized content using json package's NeWDecoder
// https://golang.org/pkg/encoding/json/#NewDecoder
func (product *Product ) FromJSON(r io.Reader) error{
	decoder := json.NewDecoder(r)
	return decoder.Decode(product)
}


// Products is a collection of Product
type Products []*Product

// ToJSON serializes contents of the collection to JSON using json package's NeWEncoder
// https://golang.org/pkg/encoding/json/#NewEncoder
func (products*Products) ToJSON(requestWriter io.Writer) error {
	encoder := json.NewEncoder(requestWriter)
	return encoder.Encode(products)
}

// GetProducts returns a list of the products
func GetProducts() Products {
	return productList
}

func AddProducts(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)
}

func UpdateProduct(id int, product*Product)  error {
	_, position, err := findProduct(id)

	if err != nil {
		return err
	}

	product.ID = id
	productList[position] = product

	return nil
}

func DeleteProduct(id int) error{
	_, position, err := findProduct(id)

	if err != nil {
		return err
	}

	productList = append(productList[:position], productList[position+1:]...)

	return nil
}

func getNextID() int {
	lastProduct := productList[len(productList) - 1]
	return lastProduct.ID + 1
}

var ErrorProductNotFound = fmt.Errorf("Product Not Found")

func findProduct(id int) (*Product, int, error) {
	for i, product := range productList {
		if product.ID == id {
			return product, i, nil
		}
	}

	return nil, -1, ErrorProductNotFound
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
