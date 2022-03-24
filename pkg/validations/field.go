package validations

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func fieldValidation(c *gin.Context, field string, validations string) map[string]string {
	var validate *validator.Validate
	err := validate.Var(field, validations)

	if err != nil {
		return map[string]string{
			field:err.Error(),
		}
	}

	return nil
}
