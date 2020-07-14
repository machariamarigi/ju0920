package handlers

import (
	"net/http"

	"github.com/machariamarigi/ju0920/product-api/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update updates a product in the data store
func (products Products) Update(responseWriter http.ResponseWriter, request *http.Request) {
	product := request.Context().Value(KeyProduct{}).(data.Product)

	products.logger.Println("[DEBUG] Handle PUT Product", product.ID)

	err := data.UpdateProduct(product)

	if err == data.ErrorProductNotFound {
		http.Error(responseWriter, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(responseWriter, "Something went wrong", http.StatusInternalServerError)
		return
	}
}
