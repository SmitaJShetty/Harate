package apperror

//AppError application error
type AppError struct {
	Message    string
	StatusCode int
}

//NewAppError returns an apperror
func NewAppError(msg string, statusCode int) *AppError {
	return &AppError{
		Message:    msg,
		StatusCode: statusCode,
	}
}
