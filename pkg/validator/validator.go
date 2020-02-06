package validator

import (
	"fmt"

	govalidator "github.com/go-playground/validator/v10"
	"github.com/supendi/orderan.api/pkg/errors"
)

//Construct a more readable error message from validation errors
func contructMessage(fieldName string, tag string, param interface{}) string {
	errorMessage := ""
	if param == nil {
		param = ""
	}
	switch tag {
	case "min":
		errorMessage = fmt.Sprintf("The minimum value of '%s' is %v", fieldName, param)
	case "max":
		errorMessage = fmt.Sprintf("The maximum value of '%s' is %v", fieldName, param)
	case "required":
		errorMessage = fmt.Sprintf("The field '%s' is required", fieldName)
	default:
		errorMessage = tag
	}
	return errorMessage
}

//Validate validates validate struct tag
func Validate(model interface{}) *errors.AppError {
	err := govalidator.New().Struct(model)
	if err == nil {
		return nil
	}
	appError := errors.NewAppError("Validation error(s) occured.")
	validationErrors := err.(govalidator.ValidationErrors)
	if validationErrors != nil {
		for _, goError := range validationErrors {
			errMessage := contructMessage(goError.StructField(), goError.Tag(), goError.Param())
			appError.Errors.Add(errors.NewFieldError(goError.StructField(), errMessage))
		}
	}
	return appError
}
