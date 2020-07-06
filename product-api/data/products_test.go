package data

import (
	"testing"
)

func TestChecksValidation(t* testing.T) {
	product := &Product{
		Name: "Mash",
		Price: 100,
		SKU: "ods-ends-inbtns",
	}

	err := product.Validate()

	if err != nil {
		t.Fatal(err)
	}
}