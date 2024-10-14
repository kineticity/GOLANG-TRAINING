package validation

import (
	"errors"
)

func ValidateNonEmptyString(fieldName, value string) error {
	if value == "" {
		return errors.New(fieldName + " cannot be empty")
	}
	return nil
}


func ValidatePositiveNumber(fieldName string, value float64) error {
	if value <= 0 {
		return errors.New(fieldName + " must be a positive number")
	}
	return nil
}

