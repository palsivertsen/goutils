package validators

import (
	"log"
	"regexp"
	"testing"
	"time"

	fuzz "github.com/google/gofuzz"
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
	validMaxLen = map[int][]string{
		1: []string{"a", "b", "c"},
		2: []string{"a", "bb"},
		3: []string{"a", "bb", "ccc"},
	}
	invalidMaxLen = map[int][]string{
		1: []string{"aa", "bbb", "cccc"},
		2: []string{"aaa", "bbbb", "ccccc"},
	}
	validMinLen = map[int][]string{
		1: []string{"a", "bb", "ccc"},
		2: []string{"aa", "bbb", "cccc"},
	}
	invalidMinLen = map[int][]string{
		1: []string{""},
		2: []string{"a", "b", "c"},
		5: []string{"aa", "bbb", "cccc"},
	}
	validLen = map[int][]string{
		1: []string{"a", "b", "c"},
		2: []string{"aa", "bb", "cc"},
	}
	invalidLen = map[int][]string{
		1: []string{"aa", "bb", "cc"},
		2: []string{"aaa", "bbb", "ccc"},
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

func TestMaxLen(t *testing.T) {
	for l, values := range validMaxLen {
		for _, value := range values {
			assert.True(t, String(value).IsMaxLen(l), "%s !max %d", value, l)
			assert.Nil(t, String(value).MaxLen(l).FirstError(), "%s !max %d", value, l)
		}
	}
	for l, values := range invalidMaxLen {
		for _, value := range values {
			assert.False(t, String(value).IsMaxLen(l), "%s !max %d", value, l)
			assert.Equal(t, ErrMax, String(value).MaxLen(l).FirstError(), "%s !max %d", value, l)
		}
	}
}

func TestMinLen(t *testing.T) {
	for l, values := range validMinLen {
		for _, value := range values {
			assert.True(t, String(value).IsMinLen(l), "%s !min %d", value, l)
			assert.Nil(t, String(value).MinLen(l).FirstError(), "%s !min %d", value, l)
		}
	}
	for l, values := range invalidMinLen {
		for _, value := range values {
			assert.False(t, String(value).IsMinLen(l), "%s !min %d", value, l)
			assert.Equal(t, ErrMin, String(value).MinLen(l).FirstError(), "%s !min %d", value, l)
		}
	}
}

func TestLen(t *testing.T) {
	for l, values := range validLen {
		for _, value := range values {
			assert.True(t, String(value).IsLen(l), "%s !len %d", value, l)
			assert.Nil(t, String(value).Len(l).FirstError(), "%s !len %d", value, l)
		}
	}
	for l, values := range invalidLen {
		for _, value := range values {
			assert.False(t, String(value).IsLen(l), "%s !len %d", value, l)
			assert.Equal(t, ErrLen, String(value).Len(l).FirstError(), "%s !len %d", value, l)
		}
	}
}

func TestFuzz(t *testing.T) {
	seed := time.Now().UnixNano()
	f := fuzz.NewWithSeed(seed)
	var value string
	f.Fuzz(&value)
	log.Printf("Seed: %d\tValue: %s", seed, value)

	unit := String(value)

	unit.Hex()
	unit.IsHex()

	unit.IsFunc(func(arg1 string) bool {
		var b bool
		f.Fuzz(&b)
		return b
	})

	var isLen int
	f.Fuzz(&isLen)
	unit.IsLen(isLen)

	var isMatch regexp.Regexp
	f.Fuzz(&isMatch)
	unit.IsMatch(isMatch.String())

	var isMaxLen int
	f.Fuzz(&isMaxLen)
	unit.IsMaxLen(isMaxLen)

	var isMinLen int
	f.Fuzz(&isMinLen)
	unit.IsMinLen(isMinLen)

	var len int
	f.Fuzz(&len)
	unit.Len(len)

	var match regexp.Regexp
	f.Fuzz(&match)
	unit.Match(match.String())

	var maxLen int
	f.Fuzz(&maxLen)
	unit.MaxLen(maxLen)

	var minLen int
	f.Fuzz(&minLen)
	unit.MinLen(minLen)
}
