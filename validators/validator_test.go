package validators

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidator(t *testing.T) {
	validator := Validator{}
	assert.False(t, validator.HasErrors())
	assert.Nil(t, validator.FirstError())
	validator.Error(errors.New("first error"))
	validator.Error(errors.New("second error"))
	assert.True(t, validator.HasErrors())
	assert.EqualError(t, validator.FirstError(), "first error")
	errs := validator.Errors()
	assert.EqualError(t, errs[0], "first error")
	assert.EqualError(t, errs[1], "second error")
}
