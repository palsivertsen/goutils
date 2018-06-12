package validators

import (
	"errors"
	"fmt"
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

func TestVerifyWithError(t *testing.T) {
	validator := Validator{}
	validator.verifyWithError(true, errors.New("this error is NOT added"))
	assert.False(t, validator.HasErrors())
	validator.verifyWithError(false, errors.New("this error is added"))
	assert.True(t, validator.HasErrors())
	assert.Error(t, validator.FirstError(), "this error is added")
}

func ExampleValidator_stringHex() {
	String("qwerty").
		Hex().
		HasErrors() // true
}

func ExampleValidator_stringIsHex() {
	String("qwerty").
		IsHex() // false
}

func ExampleValidator_chaining() {
	myValidator := String("golang"). // Create a string validator
						MaxLen(5).      // Assert maximum length
						Hex().          // Assert string is hex
						Match("[A_Z]+") // Assert string matches regex pattern

	fmt.Println(myValidator.HasErrors()) // Checks if any of the assertions failed
	fmt.Println(myValidator.Errors())    // List of assertions that failed

	// Output:
	// true
	// [String is too long String is not hex String does not match]
}
