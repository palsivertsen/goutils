package validators

import (
	"encoding/hex"
	"errors"
)

var (
	ErrNotHex = errors.New("String not hex")
)

type StringValidator struct {
	Validator
	value string
}

func String(value string) *StringValidator {
	return &StringValidator{value: value}
}

func (v *StringValidator) IsHex() bool {
	if v.value == "" {
		return false
	}
	_, err := hex.DecodeString(v.value)
	return err == nil
}

func (v *StringValidator) Hex() *StringValidator {
	if !v.IsHex() {
		v.Error(ErrNotHex)
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
