package exception

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Messages []string
}

func (e *ValidationError) Error() string {
	return "validation error"
}

func NewValidationError(validationErrors validator.ValidationErrors) *ValidationError {
	messages := make([]string, 0, len(validationErrors))
	for _, e := range validationErrors {
		messages = append(messages, fmt.Sprintf("%s: %s", e.Field(), e.Tag()))
	}
	return &ValidationError{Messages: messages}
}
