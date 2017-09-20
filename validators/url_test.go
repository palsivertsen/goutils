package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	hasScheme      = "http://localhost"
	notHasScheme   = "localhost"
	hasHostname    = "http://localhost"
	notHasHostname = "http://"
	unparsable     = ":"
)

func TestParserError(t *testing.T) {
	v := URL(unparsable)
	assert.True(t, v.HasErrors())
	assert.Error(t, v.FirstError(), ErrURLParse)
	assert.NotPanics(t, func() {
		v.HasHostname()
		v.HasScheme()
	})
}

func TestScheme(t *testing.T) {
	valid := URL(hasScheme)
	assert.True(t, valid.HasScheme())
	valid.Scheme()
	assert.False(t, valid.HasErrors())

	invalid := URL(notHasScheme)
	assert.False(t, invalid.HasScheme())
	invalid.Scheme()
	assert.True(t, invalid.HasErrors())
	assert.Error(t, invalid.FirstError(), ErrURLScheme)
}

func TestHostname(t *testing.T) {
	valid := URL(hasHostname)
	assert.True(t, valid.HasHostname())
	valid.Hostname()
	assert.False(t, valid.HasErrors())

	invalid := URL(notHasHostname)
	assert.False(t, invalid.HasHostname())
	invalid.Hostname()
	assert.True(t, invalid.HasErrors())
	assert.Error(t, invalid.FirstError(), ErrURLHostname)
}
