package validator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateMobile(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	return validateMobile(value)
}

func validateMobile(value string) bool {
	if value == "" {
		return true
	}

	var validMobile = regexp.MustCompile(`^(01[0|1])\d{7,8}$`)
	if validMobile.MatchString(value) {
		return true
	}
	return false
}
