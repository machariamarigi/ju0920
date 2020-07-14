package handlers

import (
	"net/http"

	"github.com/machariamarigi/ju0920/product-api/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// ListAll handles GET requests for all products
func (products *Products) ListAll(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("[DEBUG] Handle GET all products")

	// fetch the products from the datastore
	productsList := data.GetProducts()

	// serailize the list to JSON
	err := data.ToJSON(productsList, responseWriter)
	if err != nil {
		http.Error(responseWriter, "unable to encode json", http.StatusInternalServerError)
	}
}

// swagger:route GET /products/{id} products listSingle
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// ListSingle handles GET request for a single product
func (products *Products) ListSingle(responseWriter http.ResponseWriter, request *http.Request) {
	id := getProductID(request)

	products.logger.Println("[DEBUG] Handle GET product", id)

	product, err := data.GetProductByID(id)

	switch err {
		case nil:

		case data.ErrorProductNotFound:
			products.logger.Println("[ERROR] fetching product", err)

			responseWriter.WriteHeader(http.StatusNotFound)
			data.ToJSON(&GenericError{Message: err.Error()}, responseWriter)
			return
		
		default:
			products.logger.Println("[ERROR] fetching product", err)

			responseWriter.WriteHeader(http.StatusInternalServerError)
			data.ToJSON(&GenericError{Message: err.Error()}, responseWriter)
			return
	}

	err = data.ToJSON(product, responseWriter)

	if err != nil {
		products.logger.Println("[ERROR] serializing product", err)
	}
}
