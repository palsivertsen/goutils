package validators

type Validator struct {
	errors []error
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) != 0
}

func (v *Validator) Error(err error) {
	v.errors = append(v.errors, err)
}

func (v *Validator) Errors() []error {
	return v.errors
}

func (v *Validator) FirstError() error {
	if len(v.errors) == 0 {
		return nil
	}
	return v.errors[0]
}
