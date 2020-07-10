package data

import (
	"fmt"
	"encoding/json"
	"io"
	"time"
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Product defines the structure of an API product
type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Price        float32 `json:"price" validate:"gt=0"`
	SKU          string  `json:"sku" validate:"required,sku"`
	CreatedOn    string  `json:"_"`
	UpdatedOn    string  `json:"_"`
	DeletedOn    string  `json:"_"`
}

// ErrorProductNotFound is an error raised when a product is not found in the store
var ErrorProductNotFound = fmt.Errorf("Product Not Found")

// FromJSON decodes json serialized content using json package's NeWDecoder
// https://golang.org/pkg/encoding/json/#NewDecoder
func (product *Product) FromJSON(r io.Reader) error{
	decoder := json.NewDecoder(r)
	return decoder.Decode(product)
}

// Validate does JSON validation for our products using the Package Validator
// https://github.com/go-playground/validator
func (product *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", skuValidation)

	return validate.Struct(product)
}

func skuValidation(fl validator.FieldLevel) bool {
	// SKU format is ddfg-eews-fffr

	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regex.FindAllString(fl.Field().String(), -1)

	if len(matches) != 1 {
		return false
	}

	return true
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

// AddProduct adds a new product to the store
func AddProduct(product *Product) {
	product.ID = getNextID()
	productList = append(productList, product)
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
