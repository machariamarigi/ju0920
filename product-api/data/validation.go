package data

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
)

// ValidationError wraps validator's field error so we don't 
// expose this to out code
type ValidationError struct {
	validator.FieldError
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error: Field validation for %s failed on '%s' tag",
		validationError.Namespace(),
		validationError.Field(),
		validationError.Tag(),
	)
}

// ValidationErrors is a collection of ValidationError
type ValidationErrors []ValidationError

// Errors convers the slice into a string slice
func (validationErrors ValidationErrors) Errors() []string {
	errs := []string{}

	for _, err := range validationErrors {
		errs = append(errs, err.Error())
	}

	return errs
}

// Validation contains the validator settings and cache
type Validation struct {
	validate *validator.Validate
}

// NewValidation returns a new validation type
// registers sku validation
func NewValidation() *Validation {
	validate := validator.New()
	validate.RegisterValidation("sku", skuValidation)

	return &Validation{validate}
}

func skuValidation(fl validator.FieldLevel) bool {
	// SKU format is ddfg-eews-fffr

	regex := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := regex.FindAllString(fl.Field().String(), -1)

	if len(matches) == 1 {
		return true
	}

	return false
}

// Validate the item
// for more detail the returned error can be cast into a
// validator.ValidationErrors collection
//
// if ve, ok := err.(validator.ValidationErrors); ok {
//			fmt.Println(ve.Namespace())
//			fmt.Println(ve.Field())
//			fmt.Println(ve.StructNamespace())
//			fmt.Println(ve.StructField())
//			fmt.Println(ve.Tag())
//			fmt.Println(ve.ActualTag())
//			fmt.Println(ve.Kind())
//			fmt.Println(ve.Type())
//			fmt.Println(ve.Value())
//			fmt.Println(ve.Param())
//			fmt.Println()
//	}
func (validation *Validation) Validate(inf interface{}) ValidationErrors {
	errs := validation.validate.Struct(inf).(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	var returnErrs []ValidationError

	for _, err := range errs {
		validationError := ValidationError{err.(validator.FieldError)}
		returnErrs = append(returnErrs, validationError)
	}

	return returnErrs
}

