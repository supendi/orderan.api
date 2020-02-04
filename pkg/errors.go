package pkg

//AppError represent base known error
type AppError struct {
	message string
}

//Error return the error message
func (me *AppError) Error() string {
	return me.message
}
