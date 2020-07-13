package handlers

import (
	"net/http"

	"github.com/machariamarigi/ju0920/product-api/data"
)

// ListAll handles GET requests for all products
func (product*Products) ListAll(responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	products := data.GetProducts()

	// serailize the list to JSON
	err := data.ToJSON(products, responseWriter)
	if err != nil {
		http.Error(responseWriter, "unable to encode json", http.StatusInternalServerError)
	}
}