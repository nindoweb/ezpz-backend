package validations

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidateStruct(c *gin.Context, model interface{}) (map[string]string) {
	if err := c.ShouldBind(&model); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			return structValidation(verr)
		}
	}
	return nil
}