package validators

// Validator is the base struct for validators
// Has functions for registering errors and checking results of validations
type Validator struct {
	errors []error
}

// HasErrors checks if the validator has reported any errors
func (v *Validator) HasErrors() bool {
	return len(v.errors) != 0
}

// Error registers an error with the validator
func (v *Validator) Error(err error) {
	v.errors = append(v.errors, err)
}

// Errors returns all reported errors
func (v *Validator) Errors() []error {
	return v.errors
}

// FirstError returns the first reported error
func (v *Validator) FirstError() error {
	if len(v.errors) == 0 {
		return nil
	}
	return v.errors[0]
}
