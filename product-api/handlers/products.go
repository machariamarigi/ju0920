package handlers

import (
	"github.com/machariamarigi/ju0920/product-api/data"
	"net/http"
	"log"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p*Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// handle get
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// handlePut

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p*Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	pl := data.GetProducts()
	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to encode json", http.StatusInternalServerError)
	}
}
