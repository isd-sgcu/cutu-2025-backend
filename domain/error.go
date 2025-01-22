package domain

import "errors"

// ErrorResponse represents a basic error structure
type ErrorResponse struct {
	Error string `json:"error"`
}

var ErrUserAlreadyEntered = errors.New("user has already entered")
