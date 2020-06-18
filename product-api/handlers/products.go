package handlers

import (
	"github.com/machariamarigi/ju0920/product-api/data"
	"net/http"
	"log"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle GET
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handle POST
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}



	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	pl := data.GetProducts()

	// serailize the list to JSON
	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusInternalServerError)
	}
}

func (p*Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)

	if err != nil {
		http.Error(rw, "Unable to decode JSON", http.StatusBadRequest)
	}

	data.AddProducts(prod)
}