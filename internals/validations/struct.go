package validations

import (
	"github.com/go-playground/validator/v10"
)

func Struct(v interface{}) map[string]string {
	var validate *validator.Validate
	err := validate.Struct(v)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return nil
		}
		var validateError map[string]string
		for _, err := range err.(validator.ValidationErrors) {
			validateError[err.ActualTag()] = err.Error()
		}

		return validateError
	}

	return nil
}
