package helpers

import "github.com/go-playground/validator/v10"

func ValidationError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "This field is required"
	default:
		return fieldError.Error()
	}
}