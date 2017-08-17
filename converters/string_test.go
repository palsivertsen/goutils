package converters

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	validURLs   = []string{"http://google.com/"}
	invalidURLs = []string{":"}
)

func TestURL(t *testing.T) {
	for _, validURL := range validURLs {
		converter := String(validURL)
		u, err := converter.ToURL()
		assert.NotNil(t, u)
		assert.NoError(t, err)
		assert.NotPanics(t, func() { converter.MustURL() })
	}
	for _, invalidURL := range invalidURLs {
		converter := String(invalidURL)
		u, err := converter.ToURL()
		assert.Nil(t, u)
		assert.Error(t, err)
		assert.Panics(t, func() { converter.MustURL() })
	}
}

func TestMustTimeFromNsec(t *testing.T) {
	assert.Equal(t, time.Unix(1502993369, 783964911), String("1502993369783964911").MustTimeFromNsec())
	// TODO: Test panic
}

func TestToTimeFromNsec(t *testing.T) {
	a, err := String("1502993369783964911").ToTimeFromNsec()
	assert.Equal(t, time.Unix(1502993369, 783964911), a)
	assert.Nil(t, err)
	// TODO: Test error
}

func TestMustTimeFromSec(t *testing.T) {
	assert.Equal(t, time.Unix(1502993369, 0), String("1502993369").MustTimeFromSec())
	// TODO: Test panic
}

func TestToTimeFromSec(t *testing.T) {
	a, err := String("1502993369").ToTimeFromSec()
	assert.Equal(t, time.Unix(1502993369, 0), a)
	assert.Nil(t, err)
	// TODO: Test error
}
