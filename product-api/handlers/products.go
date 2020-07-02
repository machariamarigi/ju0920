package handlers

import (
	"strconv"
	"github.com/gorilla/mux"
	"github.com/machariamarigi/ju0920/product-api/data"
	"net/http"
	"log"
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

	product := &data.Product{}

	err := product.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to decode JSON", http.StatusBadRequest)
	}

	data.AddProducts(product)
}

// UpdateProduct updates a product in the data store
func (products Products) UpdateProduct(responseWriter http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(mux.Vars(request)["id"])

	if err != nil {
		http.Error(responseWriter, "Unable to convert id", http.StatusBadRequest)
		return
	}

	products.logger.Println("Handle PUT Product")

	product := &data.Product{}

	err = product.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to decode JSON", http.StatusBadRequest)
	}
	err = data.UpdateProduct(id, product)

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
