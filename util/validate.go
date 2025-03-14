package util

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateJSON(object any) (validates []string) {
	if err := validate.Struct(object); err != nil {
		var validateErrs validator.ValidationErrors
		if errors.As(err, &validateErrs) {
			for _, e := range validateErrs {
				validates = append(validates, e.Error())
			}
		}
	}

	return
}
