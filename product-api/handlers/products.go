package handlers

import (
	"strconv"
	"github.com/gorilla/mux"
	"github.com/machariamarigi/ju0920/product-api/data"
	"net/http"
	"log"
)

// Products handler for CRUD operations
type Products struct {
	logger     *log.Logger
	validator  *data.Validation
}

// KeyProduct is a key used for the Product object in the context
type KeyProduct struct{}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"message"`
}

// NewProducts returns a new products handler with the given logger & validator
func NewProducts(logger *log.Logger, validator *data.Validation) *Products {
	return &Products{logger, validator}
}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(request *http.Request) int {
	vars := mux.Vars(request)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
