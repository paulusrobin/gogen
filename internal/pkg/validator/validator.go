package validator

import (
	"errors"
	"sync"

	"github.com/go-playground/validator/v10"
)

const (
	defaultMessage = "validation Error"
)

var (
	v    *validator.Validate
	once sync.Once
)

// Struct validates a structs exposed fields
func Struct(s interface{}) error {
	once.Do(func() {
		v = validator.New()
	})

	if v == nil {
		return errors.New("missing validator")
	}

	err := v.Struct(s)

	if err != nil {
		e := &ValidationError{
			Message: defaultMessage,
			Details: map[string]string{},
		}

		for _, i := range err.(validator.ValidationErrors) {
			e.Details[i.StructNamespace()] = i.Error()
		}

		return e
	}
	return nil
}
