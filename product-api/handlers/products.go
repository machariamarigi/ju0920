// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import (
	"strconv"
	"context"
	"github.com/gorilla/mux"
	"github.com/machariamarigi/ju0920/product-api/data"
	"net/http"
	"log"
	"fmt"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}


// GetProducts returns products frome the data store
func (product*Products) GetProducts(responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	products := data.GetProducts()

	// serailize the list to JSON
	err := products.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "unable to encode json", http.StatusInternalServerError)
	}
}

func (products*Products) AddProduct(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle POST Products")

	product := request.Context().Value(KeyProduct{}).(data.Product)

	data.AddProducts(&product)
}

// UpdateProduct updates a product in the data store
func (products Products) UpdateProduct(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		http.Error(responseWriter, "Unable to convert id", http.StatusBadRequest)
		return
	}

	products.logger.Println("Handle PUT Product", id)

	product := request.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdateProduct(id, &product)

	if err == data.ErrorProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(responseWriter, "Something went wrong", http.StatusInternalServerError)
		return
	}
}

func (product*Products) DeleteProduct(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		http.Error(responseWriter, "Unable to convert id", http.StatusBadRequest)
		return
	}


	product.logger.Println("Handle DELETE Products")

	data.DeleteProduct(id)
}

type KeyProduct struct{}

func (products *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		product := data.Product{}

		err := product.FromJSON(request.Body)
		if err != nil {
			http.Error(responseWriter, "Unable to decode JSON", http.StatusBadRequest)
			return
		}

		err = product.Validate()
		if err != nil {
			products.logger.Println("[ERROR] validating product", err)
			http.Error(
				responseWriter,
				fmt.Sprintf("Error Validating product: %s", err),
				http.StatusBadRequest)
		}
		
		// add the product to the context
		ctx := context.WithValue(request.Context(), KeyProduct{}, product)
		request = request.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(responseWriter, request)
	})
}
