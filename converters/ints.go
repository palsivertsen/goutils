package converters

import (
	"strconv"
	"time"
)

// IntConverter is a converter for converting integers to various standard types
// All conversions are done with int64 as base
type IntConverter struct {
	i int64
}

// Int creates an new IntConverter by calling Int64()
func Int(i int) *IntConverter {
	return Int64(int64(i))
}

// Int64 creates an IntConverter
func Int64(i int64) *IntConverter {
	return &IntConverter{i: i}
}

// ToTimeFromNsec converts from nanoseconds to time.Time
// See time.Unix
func (c *IntConverter) ToTimeFromNsec() time.Time {
	return time.Unix(0, c.i)
}

// ToTimeFromSec converts from seconds to time.Time
// See time.Unix
func (c *IntConverter) ToTimeFromSec() time.Time {
	return time.Unix(c.i, 0)
}

// ToString returns a string representation of given integer
func (c *IntConverter) ToString() string {
	return strconv.FormatInt(c.i, 10)
}
