package converters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	assert.Equal(t, "2147483647", Int(2147483647).ToString())
}
