package utils

import "net/http"

// APIError Interface Generic error
type APIError interface {
	Message() string
	StatusCode() int
}

type apiError struct {
	ErrorMessage string `json:"message"`
	ErrorStatus  int    `json:"status"`
}

func (e apiError) Message() string {
	return e.ErrorMessage
}

func (e apiError) StatusCode() int {
	return e.ErrorStatus
}

// NewAPIError Creates a new apiError
func NewAPIError(message string, status int) APIError {
	return apiError{message, status}
}

// NewNotFoundAPIError Creates a new apiError with not found status
func NewNotFoundAPIError(message string) APIError {
	return apiError{message, http.StatusNotFound}
}

// NewInternalServerAPIError Creates a new apiError with internal server error status
func NewInternalServerAPIError(message string) APIError {
	return apiError{message, http.StatusInternalServerError}
}

// NewBadRequestAPIError Creates a new apiError with bad request error
func NewBadRequestAPIError(message string) APIError {
	return apiError{message, http.StatusBadRequest}
}
