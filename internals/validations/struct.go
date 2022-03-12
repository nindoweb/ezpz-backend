package validations

import (
	"ezpz/internals/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Struct(c *gin.Context, v interface{}) {
	var validate *validator.Validate
	err := validate.Struct(v)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {

		}
		var validateError map[string]string
		for _, err := range err.(validator.ValidationErrors) {
			validateError[err.ActualTag()] = err.Error()
		}

		c.JSON(http.StatusUnprocessableEntity, response.JsonResponse{
			Message: "error",
			Data:    validateError,
		})
		c.Abort()
		return
	}

}
