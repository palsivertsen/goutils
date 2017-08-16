package converters

import (
	"testing"

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
