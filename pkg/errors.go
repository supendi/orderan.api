package pkg

// FieldError Represent field error. Provides the field name which error is, and the error message
type FieldError struct {
	Field     string
	Message   string
	SubErrors []*FieldError
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

//AppError represent base known error
type AppError struct {
	ErrorMessage string
	Errors       []*FieldError
}

//NewAppError return a new app error instance
func NewAppError(errorMessage string) *AppError {
	return &AppError{
		ErrorMessage: errorMessage,
	}
}

//Error return the error message
func (me *AppError) Error() string {
	return me.ErrorMessage
}
