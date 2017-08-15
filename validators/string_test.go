package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	validHex   = []string{"abcd", "ABCD", "1234", "abcdefABCDEF1234567890"}
	invalidHex = []string{" abc", "abcdefg", "", " ", "abc"}
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
