package converters

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	validURLs     = []string{"http://google.com/"}
	invalidURLs   = []string{":"}
	validInts     = map[string]int{"-2": -2, "-1": -1, "0": 0, "1": 1, "2": 2, "09": 9}
	invalidInts   = []string{"a", "b", "c"}
	validInt8s    = map[string]int8{"-2": -2, "-1": -1, "0": 0, "1": 1, "2": 2, "09": 9}
	invalidInt8s  = invalidInts
	validInt16s   = map[string]int16{"-2": -2, "-1": -1, "0": 0, "1": 1, "2": 2, "09": 9}
	invalidInt16s = invalidInts
	validInt32s   = map[string]int32{"-2": -2, "-1": -1, "0": 0, "1": 1, "2": 2, "09": 9}
	invalidInt32s = invalidInts
	validInt64s   = map[string]int64{"-2": -2, "-1": -1, "0": 0, "1": 1, "2": 2, "09": 9}
	invalidInt64s = invalidInts
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
	assert.Panics(t, func() { String("Not an int").MustTimeFromNsec() })
}

func TestToTimeFromNsec(t *testing.T) {
	{
		a, err := String("1502993369783964911").ToTimeFromNsec()
		assert.Equal(t, time.Unix(1502993369, 783964911), a)
		assert.Nil(t, err)
	}
	{
		_, err := String("Not an int").ToTimeFromNsec()
		assert.Error(t, err)
	}
}

func TestMustTimeFromSec(t *testing.T) {
	assert.Equal(t, time.Unix(1502993369, 0), String("1502993369").MustTimeFromSec())
	assert.Panics(t, func() { String("Not an int").MustTimeFromSec() })
}

func TestToTimeFromSec(t *testing.T) {
	{
		a, err := String("1502993369").ToTimeFromSec()
		assert.Equal(t, time.Unix(1502993369, 0), a)
		assert.Nil(t, err)
	}
	{
		_, err := String("Not an int").ToTimeFromSec()
		assert.Error(t, err)
	}
}

func TestTimeFormatUtcDDMMYYYYSlash(t *testing.T) {
	valids := []map[string]time.Time{{"01/02/2000": time.Date(2000, 02, 01, 0, 0, 0, 0, time.UTC)}}
	invalids := []string{" ", "not a date", "01.02.2000", "1/2/2000"}
	for _, valid := range valids {
		for value, expected := range valid {
			converter := String(value)
			actual, err := converter.ToTimeFormatUtcDDMMYYYYSlash()
			assert.Equal(t, expected, actual)
			assert.NoError(t, err, value)
			assert.NotPanics(t, func() {
				assert.Equal(t, expected, converter.MustTimeFormatUtcDDMMYYYYSlash(), value)
			}, value)
		}
	}
	for _, value := range invalids {
		converter := String(value)
		_, err := converter.ToTimeFormatUtcDDMMYYYYSlash()
		assert.Error(t, err, value)
		assert.Panics(t, func() { converter.MustTimeFormatUtcDDMMYYYYSlash() }, value)
	}
}

func TestInt(t *testing.T) {
	for i, expected := range validInts {
		converter := String(i)
		a, err := converter.ToInt()
		assert.Equal(t, expected, a, i)
		assert.NoError(t, err, i)
		assert.NotPanics(t, func() {
			assert.Equal(t, expected, converter.MustInt(), i)
		}, i)
	}
	for _, i := range invalidInts {
		converter := String(i)
		_, err := converter.ToInt()
		assert.Error(t, err, i)
		assert.Panics(t, func() { converter.MustInt() }, i)
	}
}

func TestInt8(t *testing.T) {
	for i, expected := range validInt8s {
		converter := String(i)
		a, err := converter.ToInt8()
		assert.Equal(t, expected, a, i)
		assert.NoError(t, err, i)
		assert.NotPanics(t, func() {
			assert.Equal(t, expected, converter.MustInt8(), i)
		}, i)
	}
	for _, i := range invalidInt8s {
		converter := String(i)
		_, err := converter.ToInt8()
		assert.Error(t, err, i)
		assert.Panics(t, func() { converter.MustInt8() }, i)
	}
}

func TestInt16(t *testing.T) {
	for i, expected := range validInt16s {
		converter := String(i)
		a, err := converter.ToInt16()
		assert.Equal(t, expected, a, i)
		assert.NoError(t, err, i)
		assert.NotPanics(t, func() {
			assert.Equal(t, expected, converter.MustInt16(), i)
		}, i)
	}
	for _, i := range invalidInt16s {
		converter := String(i)
		_, err := converter.ToInt16()
		assert.Error(t, err, i)
		assert.Panics(t, func() { converter.MustInt16() }, i)
	}
}

func TestInt32(t *testing.T) {
	for i, expected := range validInt32s {
		converter := String(i)
		a, err := converter.ToInt32()
		assert.Equal(t, expected, a, i)
		assert.NoError(t, err, i)
		assert.NotPanics(t, func() {
			assert.Equal(t, expected, converter.MustInt32(), i)
		}, i)
	}
	for _, i := range invalidInt32s {
		converter := String(i)
		_, err := converter.ToInt32()
		assert.Error(t, err, i)
		assert.Panics(t, func() { converter.MustInt32() }, i)
	}
}

func TestInt64(t *testing.T) {
	for i, expected := range validInt64s {
		converter := String(i)
		a, err := converter.ToInt64()
		assert.Equal(t, expected, a, i)
		assert.NoError(t, err, i)
		assert.NotPanics(t, func() {
			assert.Equal(t, expected, converter.MustInt64(), i)
		}, i)
	}
	for _, i := range invalidInt64s {
		converter := String(i)
		_, err := converter.ToInt64()
		assert.Error(t, err, i)
		assert.Panics(t, func() { converter.MustInt64() }, i)
	}
}

func ExampleStringConverter_MustTimeFromNsec() {
	t := String("1502835623824559499").MustTimeFromNsec()
	fmt.Println(t.UTC().Format(time.UnixDate))
	// Output:
	// Tue Aug 15 22:20:23 UTC 2017
}

func ExampleStringConverter_MustURL() {
	t := String("https://example.com").MustURL()
	fmt.Println(t)
	// Output:
	// https://example.com
}
