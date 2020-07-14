package handlers

import (
	"net/http"

	"github.com/machariamarigi/ju0920/product-api/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the store
func (products *Products) Delete(responseWriter http.ResponseWriter, request *http.Request) {
	id := getProductID(request)

	products.logger.Println("[DEBUG] deleting product with id ", id)

	err := data.DeleteProduct(id)
	if err == data.ErrorProductNotFound {
		products.logger.Println("[ERROR] deleting PRODUCT id does not exist")

		responseWriter.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, responseWriter)
		return
	}

	if err != nil {
		products.logger.Println("[ERROR] deleting product", err)

		responseWriter.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, responseWriter)
		return
	}

	responseWriter.WriteHeader(http.StatusNoContent)
}
