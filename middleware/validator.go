package middleware

import (
	val "github.com/YooGenie/daily-work-log-service/common/validator"
	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) (err error) {
	return cv.validator.Struct(i)
}

func RegisterValidator() *CustomValidator {
	customValidator := validator.New()
	customValidator.RegisterValidation("date8", val.ValidateDate8)
	customValidator.RegisterValidation("date12", val.ValidateDate12)
	return &CustomValidator{validator: customValidator}
}
