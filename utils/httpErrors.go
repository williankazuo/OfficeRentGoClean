package utils

// ErrorMessage Struct for Error message
type ErrorMessage struct {
	Message string `json:"message"`
}

// NewErrorMessage Creates a new error message
func NewErrorMessage(message string) ErrorMessage {
	return ErrorMessage{Message: message}
}
