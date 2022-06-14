package validator

import (
	"github.com/go-playground/validator/v10"
	"time"
)

func ValidateDate8(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return validateDate8(value)
}

func validateDate8(value string) bool {
	if value == "" {
		return true
	}

	_, err := time.Parse("20060102", value)
	if err != nil {
		return false
	}
	return true
}

func ValidateDate12(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return validateDate12(value)
}

func validateDate12(value string) bool {
	if value == "" {
		return true
	}

	_, err := time.Parse("200601021504", value)
	if err != nil {
		return false
	}
	return true
}
