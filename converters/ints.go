package converters

import (
	"strconv"
	"time"
)

type IntConverter struct {
	i int64
}

func Int(i int) *IntConverter {
	return &IntConverter{i: int64(i)}
}

func (c *IntConverter) ToTimeFromNsec() time.Time {
	return time.Unix(0, c.i)
}

func (c *IntConverter) ToTimeFromSec() time.Time {
	return time.Unix(c.i, 0)
}

func (c *IntConverter) ToString() string {
	return strconv.FormatInt(c.i, 10)
}
