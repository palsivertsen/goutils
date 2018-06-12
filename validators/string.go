package validators

import (
	"encoding/hex"
	"errors"
	"regexp"
)

var (
	// ErrNotHex indicates a Hex error
	ErrNotHex = errors.New("String is not hex")
	// ErrNoMatch indicates a Match error
	ErrNoMatch = errors.New("String does not match")
	// ErrFunc indicates an error when validating using func
	ErrFunc = errors.New("func did not validate to true")
	// ErrMax validator length too long
	ErrMax = errors.New("String is too long")
	// ErrMin validator length too short
	ErrMin = errors.New("String is too short")
	// ErrLen invalid validator length
	ErrLen = errors.New("String is not correct length")
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

// IsFunc executes the given func with the validator value as parameter and returns the result
func (v *StringValidator) IsFunc(f func(string) bool) bool {
	return f(v.value)
}

// Func registers an error with the validator if IsFunc returns false
// Is chainable
func (v *StringValidator) Func(f func(string) bool) *StringValidator {
	if !v.IsFunc(f) {
		v.Error(ErrFunc)
	}
	return v
}

// IsMaxLen checks if the validator has a maximun length of l
func (v *StringValidator) IsMaxLen(l int) bool {
	return len(v.value) <= l
}

// MaxLen registers an error with the validator if IsMaxLen returns false
// Is chainable
func (v *StringValidator) MaxLen(l int) *StringValidator {
	if !v.IsMaxLen(l) {
		v.Error(ErrMax)
	}
	return v
}

// IsMinLen checks if the validator has a minimum length of l
func (v *StringValidator) IsMinLen(l int) bool {
	return len(v.value) >= l
}

// MinLen registers an error with the validator if IsMinLen returns false
// Is chainable
func (v *StringValidator) MinLen(l int) *StringValidator {
	if !v.IsMinLen(l) {
		v.Error(ErrMin)
	}
	return v
}

// IsLen checks if the validator has a length of l
func (v *StringValidator) IsLen(l int) bool {
	return len(v.value) == l
}

// Len registers an error with the validator if IsLen returns false
// Is chainable
func (v *StringValidator) Len(l int) *StringValidator {
	if !v.IsLen(l) {
		v.Error(ErrLen)
	}
	return v
}
