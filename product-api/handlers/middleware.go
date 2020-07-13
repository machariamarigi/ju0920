package handlers

import (
	"context"
	"net/http"

	"github.com/machariamarigi/ju0920/product-api/data"
)

// MiddlewareProductValidation validates the product in the request and calls next if ok
func (products *Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		product := &data.Product{}

		err := data.FromJSON(product, request.Body)
		if err != nil {
			products.logger.Println("[ERROR] deserializing product", err)
			responseWriter.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, responseWriter)
			return
		}

		errs := products.validator.Validate(product)
		if len(errs) != 0 {
			products.logger.Println("[ERROR] validating product", errs)

			responseWriter.WriteHeader(http.StatusUnprocessableEntity)
			data.ToJSON(&ValidationError{Messages: errs.Errors()}, responseWriter)
			return
		}
		
		// add the product to the context
		ctx := context.WithValue(request.Context(), KeyProduct{}, product)
		request = request.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(responseWriter, request)
	})
}
