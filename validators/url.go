package validators

import (
	"errors"
	"net/url"
)

var (
	// ErrURLParse indicates there was an error parsing the given string
	ErrURLParse = errors.New("Could not parse string")
	// ErrURLHostname indicates there was no hostname
	ErrURLHostname = errors.New("No hostname")
	// ErrURLScheme indicates there was no scheme
	ErrURLScheme = errors.New("No scheme")
)

// URLValidator is a validator used to validate urls
// Validations can be chained. See Validator for information on how to check errors
// Use IsXX functions to check a single validation
type URLValidator struct {
	Validator
	value url.URL
}

// URL creates a new URLValidator
// Takes a string to validate
func URL(value string) *URLValidator {
	u, err := url.Parse(value)
	if err != nil {
		validator := &URLValidator{value: url.URL{}}
		validator.Error(ErrURLParse)
		return validator
	}
	return &URLValidator{value: *u}
}

// HasScheme checks if url scheme != ""
func (v *URLValidator) HasScheme() bool {
	return v.value.Scheme != ""
}

// Scheme registers an error with the validator if HasScheme returns false
// Is chainable
func (v *URLValidator) Scheme() *URLValidator {
	v.verifyWithError(v.HasScheme(), ErrURLScheme)
	return v
}

// HasHostname checks if url.Hostname != ""
func (v *URLValidator) HasHostname() bool {
	return v.value.Hostname() != ""
}

// Hostname registers an ErrURLHostname error with the validator if HasHostanme returns false
// Is chainable
func (v *URLValidator) Hostname() *URLValidator {
	v.verifyWithError(v.HasHostname(), ErrURLHostname)
	return v
}
