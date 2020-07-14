package handlers

import (
	"net/http"

	"github.com/machariamarigi/ju0920/product-api/data"
)

// swagger:route POST /products products createProduct
// Create a new product
//
// responses:
//	200: productResponse
// 	422: errorValidation
// 	501: errorResponse

// Create handles POST requests for adding a new product
func (products*Products) Create(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle POST Products")

	product := request.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(product)
}
