package validation

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ivmello/finanzen-api/internal/adapter/api/exception"
)

var validate = validator.New()

func ValidateStruct(inputDto interface{}) *exception.ErrorResponse {
	var errors []*exception.ErrorCauses
	err := validate.Struct(inputDto)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element exception.ErrorCauses
			element.Field = err.StructField()
			element.Message = msgForTag(err.Tag(), err.Param())
			errors = append(errors, &element)
		}
	}
	errMessage := exception.NewException("Invalid validation on fields", 400, errors)
	if errMessage.Causes != nil {
		return errMessage
	}
	return nil
}

func msgForTag(tag string, value string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("Required min size is %s", value)
	case "max":
		return fmt.Sprintf("Required max size is %s", value)
	}
	return ""
}
