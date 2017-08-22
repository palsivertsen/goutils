package converters

import (
	"net/url"
	"strconv"
	"time"
)

// StringConverter is a converter for converting strings to various standard types
type StringConverter struct {
	from string
}

// String creates an new StringConverter
func String(from string) *StringConverter {
	return &StringConverter{from: from}
}

// ToTimeFromNsec converts from nanoseconds to time.Time
// Returns error if string is not an int
func (c *StringConverter) ToTimeFromNsec() (time.Time, error) {
	i, err := c.ToInt64()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, i), nil
}

// MustTimeFromNsec calls ToTimeFromNsec and panics on error
func (c *StringConverter) MustTimeFromNsec() time.Time {
	t, err := c.ToTimeFromNsec()
	panicOnError(err)
	return t
}

// ToTimeFromSec converts from seconds to time.Time
// Returns error if string is not an int
func (c *StringConverter) ToTimeFromSec() (time.Time, error) {
	i, err := c.ToInt64()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

// MustTimeFromSec calls ToTimeFromSec and panics on error
func (c *StringConverter) MustTimeFromSec() time.Time {
	t, err := c.ToTimeFromSec()
	panicOnError(err)
	return t
}

// ToTimeFormatUtcDDMMYYYYSlash expects a string in "DD/MM/YY" format and converts it to time.Time
func (c *StringConverter) ToTimeFormatUtcDDMMYYYYSlash() (time.Time, error) {
	return time.Parse("02/01/2006", c.from)
}

// MustTimeFormatUtcDDMMYYYYSlash calls ToTimeFormatUtcDDMMYYYYSlash and panics on error
func (c *StringConverter) MustTimeFormatUtcDDMMYYYYSlash() time.Time {
	t, err := c.ToTimeFormatUtcDDMMYYYYSlash()
	if err != nil {
		panic(err)
	}
	return t
}

// ToInt converts to int. Returns error if string is not an int
func (c *StringConverter) ToInt() (int, error) {
	i, err := c.toInt(0)
	return int(i), err
}

// MustInt calls ToInt and panics on error
func (c *StringConverter) MustInt() int {
	i, err := c.ToInt()
	panicOnError(err)
	return i
}

// ToInt8 converts to int8. Returns error if string is not an int8
func (c *StringConverter) ToInt8() (int8, error) {
	i, err := c.toInt(0)
	return int8(i), err
}

// MustInt8 calls ToInt8 and panics on error
func (c *StringConverter) MustInt8() int8 {
	i, err := c.ToInt8()
	panicOnError(err)
	return i
}

// ToInt16 converts to int16. Returns error if string is not an int16
func (c *StringConverter) ToInt16() (int16, error) {
	i, err := c.toInt(0)
	return int16(i), err
}

// MustInt16 calls ToInt16 and panics on error
func (c *StringConverter) MustInt16() int16 {
	i, err := c.ToInt16()
	panicOnError(err)
	return i
}

// ToInt32 converts to int32. Returns error if string is not an int32
func (c *StringConverter) ToInt32() (int32, error) {
	i, err := c.toInt(0)
	return int32(i), err
}

// MustInt32 calls ToInt32 and panics on error
func (c *StringConverter) MustInt32() int32 {
	i, err := c.ToInt32()
	panicOnError(err)
	return i
}

// ToInt64 converts to int64. Returns error if string is not an int64
func (c *StringConverter) ToInt64() (int64, error) {
	return c.toInt(64)
}

// MustInt64 calls ToInt64 and panics on error
func (c *StringConverter) MustInt64() int64 {
	i, err := c.ToInt64()
	panicOnError(err)
	return i
}

// ToURL converts to url.URL. Returns error on syntax errors.
// See url.Parse
func (c *StringConverter) ToURL() (*url.URL, error) {
	return url.Parse(c.from)
}

// MustURL calls ToURL and panics on error
func (c *StringConverter) MustURL() *url.URL {
	u, err := c.ToURL()
	panicOnError(err)
	return u
}

func (c *StringConverter) toInt(bitSize int) (int64, error) {
	i, err := strconv.ParseInt(c.from, 0, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
