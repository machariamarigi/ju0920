package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Hello is a simple handeler
type Hello struct {
	l *log.Logger
}

// NewHello creates a new Hello handler with the given logger
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// ServeHTTP implements the go http.Handler interface
// https://golang.org/pkg/net/http/#Handler
func (h* Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World");

	// read the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)
		http.Error(rw, "Oops", http.StatusBadRequest)
		return
	}

	// write the response
	fmt.Fprintf(rw, "Hello %s", d)
}
