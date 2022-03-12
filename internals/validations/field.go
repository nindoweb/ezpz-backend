package validations

import (
	"ezpz/internals/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func Field(c *gin.Context, field string, validations string) {
	var validate *validator.Validate
	err := validate.Var(field, validations)

	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, response.JsonResponse{
			Message: "error",
			Data: map[string]string{
				field: err.Error(),
			},
		})
		c.Abort()
		return
	}
}
