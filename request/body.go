package request

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled())
)

// Tries to extract the struct of type T from the JSON body of the request, and validates
// this struct using the validator library
//
// Most commonly used with DTO structs that have validation from the validator library
//
//	type SumDTO struct {
//		First  int `json:"first" validate:"required"`
//		Second int `json:"second" validate:"required,gt=0,lt=50"`
//	}
//
//	func SumHandler() http.HandlerFunc {
//		return func(w http.ResponseWriter, r *http.Request) {
//			dto, err := request.DecodeAndValidateJSON[SumDTO](r)
//
//			if err != nil {
//				response.To(w).Err(httperr.Bad("Invalid SumDTO")).SendJSON()
//				return
//			}
//
//			response.To(w).Content(dto.First + dto.Second)
//		}
//	}
func DecodeAndValidateJSON[T any](r *http.Request) (*T, error) {
	var target T

	err := json.NewDecoder(r.Body).Decode(target)

	if err != nil {
		return nil, err
	}

	err = validate.Struct(target)

	if err != nil {
		return nil, err
	}

	return &target, nil
}
