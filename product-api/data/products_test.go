package data

import (
	"testing"
	"bytes"

	"github.com/stretchr/testify/assert"
)

func TestProductMissingNameReturnsError(testing *testing.T) {
	product := Product{
		Price: 1.22,
		SKU:   "dffg-fgfg-dfdf",
	}

	validation := NewValidation()
	err := validation.Validate(product)
	assert.Len(testing, err, 1)
}

func TestProductMissingPriceReturnsError(testing *testing.T) {
	product := Product{
		Name:  "test product",
		Price: -1,
		SKU:   "dffg-fgfg-dfdf",
	}

	validation := NewValidation()
	err := validation.Validate(product)
	assert.Len(testing, err, 1)
}

func TestInvalidSKUReturnsError(testing *testing.T) {
	product := Product{
		Name:  "test product",
		Price: 1.22,
		SKU:   "dffg",
	}
	validation := NewValidation()
	err := validation.Validate(product)
	assert.Len(testing, err, 1)
}

func TestValidProductDoesNotReturnError(testing *testing.T) {
	product := Product{
		Name:  "test product",
		Price: 1.23,
		SKU:   "dffg-fgfg-dfdf",
	}

	validation := NewValidation()
	err := validation.Validate(product)
	assert.Len(testing, err, 0)
}

func TestProductToJSON(testing *testing.T) {
	productSample := []*Product{
		&Product{
			Name: "test",
		},
	}

	buffer := bytes.NewBufferString("")
	err := ToJSON(productSample, buffer)
	assert.NoError(testing, err)
}
