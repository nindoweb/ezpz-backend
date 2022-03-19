package validations

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func structValidation(err validator.ValidationErrors) map[string]string {
	errs := make(map[string]string)

	for _, f := range err {
		err := f.ActualTag()
		if f.Param() != "" {
			err = fmt.Sprintf("%s=%s", err, f.Param())
		}
		errs[f.Field()] = err
	}

	return errs
}
