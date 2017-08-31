package validators

import (
	"errors"
)

var (
	// ErrMin indicates a Min error
	ErrMin = errors.New("Integer is less than Min arg")
)

// IntValidator is a validator used to validate ints
// Validations can be chained. See Validator for information on how to check errors
// Use IsXX functions to check a single validation
type IntValidator struct {
	Validator
	value int64
}

// Int creates a new IntValidator
// Value is converted to a int64
func Int(value int) *IntValidator {
	return &IntValidator{value: int64(value)}
}

// Int8 creates a new IntValidator
// Value is converted to a int64
func Int8(value int8) *IntValidator {
	return &IntValidator{value: int64(value)}
}

// Int16 creates a new IntValidator
// Value is converted to a int64
func Int16(value int16) *IntValidator {
	return &IntValidator{value: int64(value)}
}

// Int32 creates a new IntValidator
// Value is converted to a int64
func Int32(value int32) *IntValidator {
	return &IntValidator{value: int64(value)}
}

// Int64 creates a new IntValidator
// Value is converted to a int64
func Int64(value int64) *IntValidator {
	return &IntValidator{value: value}
}

// IsMin checks if value is greater than or equal to given arg
func (v *IntValidator) IsMin(min int64) bool {
	return v.value >= min
}

// Min registers an error with the validator if IsMin returns false
// Is chainable
func (v *StringValidator) Min(min int64) *StringValidator {
	if !v.IsMin(min) {
		v.Error(ErrMin)
	}
	return v
}
