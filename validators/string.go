package validators

import (
	"encoding/hex"
	"errors"
	"regexp"
)

var (
	// ErrNotHex indicates a Hex error
	ErrNotHex = errors.New("String not hex")
	// ErrNoMatch indicates a Match error
	ErrNoMatch = errors.New("String no match")
)

// StringValidator is a validator used to validate strings
// Validations can be chained. See Validator for information on how to check errors
// Use IsXX functions to check a single validation
type StringValidator struct {
	Validator
	value string
}

// String creates a new StringValidator
// Takes a string to validate
func String(value string) *StringValidator {
	return &StringValidator{value: value}
}

// IsHex checks if validator is a valid HEX string
// Rules are:
// 	[0-9a-fA-F]
// 	Even length
// 	Not zero length
func (v *StringValidator) IsHex() bool {
	if v.value == "" {
		return false
	}
	_, err := hex.DecodeString(v.value)
	return err == nil
}

// Hex registers an error with the validator if IsHex returns false
// Is chainable
func (v *StringValidator) Hex() *StringValidator {
	if !v.IsHex() {
		v.Error(ErrNotHex)
	}
	return v
}

// IsMatch checks if validator matches the given regex pattern
func (v *StringValidator) IsMatch(pattern string) bool {
	return regexp.MustCompile(pattern).MatchString(v.value)
}

// Match registers an error with the validator if IsMatch returns false
// Is chainable
func (v *StringValidator) Match(pattern string) *StringValidator {
	if !v.IsMatch(pattern) {
		v.Error(ErrNoMatch)
	}
	return v
}

//
// func (v *StringValidator) IsMaxLen(len int) *StringValidator {
//
// }
//
// func (v *StringValidator) IsMinLen(len int) *StringValidator {
//
// }
//
// func (v *StringValidator) IsMaxLen(len int) *StringValidator {
//
// }
