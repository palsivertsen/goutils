package conversion

import (
	"strconv"
	"time"
)

type StringConverter struct {
	from string
}

func String(from string) *StringConverter {
	return &StringConverter{from: from}
}

func (c *StringConverter) ToTimeFromNsec() (time.Time, error) {
	i, err := c.ToInt64()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, i), nil
}

func (c *StringConverter) ToTimeFromSec() (time.Time, error) {
	i, err := c.ToInt64()
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(i, 0), nil
}

func (c *StringConverter) ToInt() (int, error) {
	i, err := c.toInt(0)
	return int(i), err
}
func (c *StringConverter) ToInt8() (int8, error) {
	i, err := c.toInt(0)
	return int8(i), err
}
func (c *StringConverter) ToInt16() (int16, error) {
	i, err := c.toInt(0)
	return int16(i), err
}
func (c *StringConverter) ToInt32() (int32, error) {
	i, err := c.toInt(0)
	return int32(i), err
}

func (c *StringConverter) ToInt64() (int64, error) {
	return c.toInt(64)
}

func (c *StringConverter) toInt(bitSize int) (int64, error) {
	i, err := strconv.ParseInt(c.from, 0, 64)
	if err != nil {
		return 0, err
	}
	return i, nil
}
