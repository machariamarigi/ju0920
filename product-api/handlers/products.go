package handlers

import (
	"github.com/machariamarigi/ju0920/product-api/data"
	"net/http"
	"log"
	"regexp"
	"strconv"
)

type Products struct {
	logger *log.Logger
}

func NewProducts(logger *log.Logger) *Products {
	return &Products{logger}
}

func (product*Products) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	// handle GET
	if request.Method == http.MethodGet {
		product.getProducts(responseWriter, request)
		return
	}

	// handle POST
	if request.Method == http.MethodPost {
		product.addProduct(responseWriter, request)
		return
	}

	// handle PUT
	if request.Method == http.MethodPut {
		product.logger.Println("PUT", request.URL.Path)
	
		regex := regexp.MustCompile(`/([0-9]+)`)
		paramGroup := regex.FindAllStringSubmatch(request.URL.Path, -1)

		if len(paramGroup) != 1 {
			product.logger.Println("Invalid URI more than one id")
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(paramGroup[0]) != 2 {
			product.logger.Println("Invalid URI more than one capture group")
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := paramGroup[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			product.logger.Println("Invalid URI unable to convert to numer", idString)
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		product.updateProduct(id, responseWriter, request)
		return
	}

	if request.Method == http.MethodDelete {
		regex := regexp.MustCompile(`/([0-9]+)`)
		paramGroup := regex.FindAllStringSubmatch(request.URL.Path, -1)

		if len(paramGroup) != 1 {
			product.logger.Println("Invalid URI more than one id")
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		if len(paramGroup[0]) != 2 {
			product.logger.Println("Invalid URI more than one capture group")
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}

		idString := paramGroup[0][1]
		id, err := strconv.Atoi(idString)

		if err != nil {
			product.logger.Println("Invalid URI unable to convert to numer", idString)
			http.Error(responseWriter, "Invalid URL", http.StatusBadRequest)
			return
		}
		product.deleteProduct(id, responseWriter, request)
		return
	}


	// catch all
	responseWriter.WriteHeader(http.StatusMethodNotAllowed)
}

func (product*Products) getProducts(responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle GET Products")

	// fetch the products from the datastore
	products := data.GetProducts()

	// serailize the list to JSON
	err := products.ToJSON(responseWriter)
	if err != nil {
		http.Error(responseWriter, "unable to encode json", http.StatusInternalServerError)
	}
}

func (products*Products) addProduct(responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle POST Products")

	product := &data.Product{}

	err := product.FromJSON(request.Body)
	if err != nil {
		http.Error(responseWriter, "Unable to decode JSON", http.StatusBadRequest)
	}

	data.AddProducts(product)
}

func (products Products) updateProduct(id int, responseWriter http.ResponseWriter, request *http.Request) {
	products.logger.Println("Handle PUT Product")

	product := &data.Product{}

	err := product.FromJSON(request.Body)
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

func (product*Products) deleteProduct(id int, responseWriter http.ResponseWriter, request *http.Request) {
	product.logger.Println("Handle DELETE Products")

	data.DeleteProduct(id)
}
