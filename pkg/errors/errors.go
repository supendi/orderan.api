package errors

import (
	"reflect"
	"strings"
)

// FieldError Represent field error. Provides the field name which error is, and the error message
type FieldError struct {
	Field     string        `json:"field"`
	Message   string        `json:"message"`
	SubErrors []*FieldError `json:"subErrors"`
}

//NewFieldError returns a new field error instance
func NewFieldError(field string, message string) *FieldError {
	return &FieldError{
		Field:     field,
		Message:   message,
		SubErrors: []*FieldError{},
	}
}

//AddSubError add new error into sub errors
func (me *FieldError) AddSubError(fieldError *FieldError) {
	me.SubErrors = append(me.SubErrors, fieldError)
}

//FieldErrors represent a multiple field error
type FieldErrors []*FieldError

//Add add new error into sub errors
func (me *FieldErrors) Add(fieldError *FieldError) {
	*me = append(*me, fieldError)
}

//AppError represent base known error
type AppError struct {
	Message string      `json:"message"`
	Errors  FieldErrors `json:"errors"`
}

//NewAppError return a new app error instance
func NewAppError(errorMessage string) *AppError {
	return &AppError{
		Message: errorMessage,
		Errors:  []*FieldError{},
	}
}

//Error return the error message
func (me *AppError) Error() string {
	return me.Message
}

//IsAppError checks if an error type is AppError
func IsAppError(err error) bool {
	t := reflect.TypeOf(err).String()
	t = strings.ReplaceAll(t, "*", "") //just ignore pointer type
	typeName := "errors.AppError"
	return t == typeName
}

//TypeName return the error type name string
func TypeName(err error) string {
	t := reflect.TypeOf(err).String()
	return t
}
