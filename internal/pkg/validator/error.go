package validator

import (
	"fmt"
	"github.com/pkg/errors"
)

// ValidationError is validation error object contains message and map details
type ValidationError struct {
	Message string
	Details map[string]string
}

// ErrDetail is error detail object contains Field and Message
type ErrDetail struct {
	Field   string
	Message string
}

// Error function implement error interface
func (e ValidationError) Error() string {
	return e.Message
}

// Detail function to get error detail in array of string
func (e ValidationError) Detail() []string {
	details := make([]string, 0)

	for _, detail := range e.Details {
		details = append(details, detail)
	}

	return details
}

// ErrorWithDetails function to get error message in string
func (e ValidationError) ErrorWithDetails() string {
	details := ""

	for k, v := range e.Details {
		details += fmt.Sprintf(" # field '%s' (%s).", k, v)
	}

	return fmt.Sprintf("%s%s", e.Message, details)
}

// IsValidationError function to check in error is instance of ValidationError or not
// param err error to be checked
// return the cast err and true if it is the ValidationError instance
// otherwise return false if it is not a ValidationError
func IsValidationError(err error) (*ValidationError, bool) {
	var validationErr *ValidationError
	if errors.As(err, &validationErr) {
		return validationErr, true
	}
	return nil, false
}
