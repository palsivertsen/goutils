package converters

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	assert.Equal(t, "2147483647", Int(2147483647).ToString())
}

func TestDate(t *testing.T) {
	assert.Equal(t, time.Unix(0, 1502883134757394503), Int(1502883134757394503).ToTimeFromNsec())
	assert.Equal(t, time.Unix(1502883134, 0), Int(1502883134).ToTimeFromSec())
}
