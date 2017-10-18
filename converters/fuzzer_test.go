package converters

import (
	"log"
	"testing"
	"time"

	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/assert"
)

func TestFuzzIntStringInt(t *testing.T) {
	seed := time.Now().UnixNano()
	f := fuzz.NewWithSeed(seed)
	var value int
	f.Fuzz(&value)
	log.Printf("Seed: %d\tValue: %d", seed, value)

	assert.Equal(t, value, String(Int(value).ToString()).MustInt())
}

func TestFuzzInt64StringInt64(t *testing.T) {
	seed := time.Now().UnixNano()
	f := fuzz.NewWithSeed(seed)
	var value int64
	f.Fuzz(&value)
	log.Printf("Seed: %d\tValue: %d", seed, value)

	assert.Equal(t, value, String(Int64(value).ToString()).MustInt64())
}
