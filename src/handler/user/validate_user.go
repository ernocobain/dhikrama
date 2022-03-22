package user

import (
	user "dhikrama.com/web/src/model/entity/users"
	"github.com/go-playground/validator"
)

var validate = validator.New()

func ValidateUser(u user.UserRequest) []*user.ErrorResponse {
	var errors []*user.ErrorResponse
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element user.ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
