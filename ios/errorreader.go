package ios

import (
	"errors"
	"io"
)

type errorReader struct {
	text string
}

// NewErrorReader returns an io.Reader that always returns an error
func NewErrorReader(errorText string) io.Reader {
	return &errorReader{text: errorText}
}

func (e *errorReader) Read(p []byte) (int, error) {
	return 0, errors.New(e.text)
}
