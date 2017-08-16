package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validHex   = []string{"abcd", "ABCD", "1234", "abcdefABCDEF1234567890"}
	invalidHex = []string{" abc", "abcdefg", "", " ", "abc"}
	validRegex = map[string][]string{
		"^[a-z]*$": []string{"asd", "abc"},
	}
	invalidRegex = map[string][]string{
		"^[a-c]*$": []string{"abcd", "ABC"},
	}
	validFunc = map[string][]func(string) bool{
		"some string": []func(string) bool{
			func(string) bool { return true },
			func(value string) bool { return value == "some string" },
		},
	}
	invalidFunc = map[string][]func(string) bool{
		"some string": []func(string) bool{
			func(string) bool { return false },
			func(value string) bool { return value != "some string" },
		},
	}
)

func TestHex(t *testing.T) {
	for _, valid := range validHex {
		assert.True(t, String(valid).IsHex(), valid)
		assert.Nil(t, String(valid).Hex().FirstError(), valid)
	}
	for _, invalid := range invalidHex {
		assert.False(t, String(invalid).IsHex(), invalid)
		assert.Equal(t, ErrNotHex, String(invalid).Hex().FirstError(), invalid)
	}
}

func TestRegex(t *testing.T) {
	for pattern, values := range validRegex {
		for _, value := range values {
			assert.True(t, String(value).IsMatch(pattern), value)
			assert.Nil(t, String(value).Match(pattern).FirstError(), value)
		}
	}
	for pattern, values := range invalidRegex {
		for _, value := range values {
			assert.False(t, String(value).IsMatch(pattern), value)
			assert.Equal(t, ErrNoMatch, String(value).Match(pattern).FirstError(), value)
		}
	}
}

func TestFunc(t *testing.T) {
	for value, funcs := range validFunc {
		for _, f := range funcs {
			assert.True(t, String(value).IsFunc(f), value)
			assert.Nil(t, String(value).Func(f).FirstError(), value)
		}
	}
	for value, funcs := range invalidFunc {
		for _, f := range funcs {
			assert.False(t, String(value).IsFunc(f), value)
			assert.Equal(t, ErrFunc, String(value).Func(f).FirstError(), value)
		}
	}
}
