package converters

import "strconv"

type IntConverter struct {
	i int64
}

func Int(i int) *IntConverter {
	return &IntConverter{i: int64(i)}
}

func (c *IntConverter) ToString() string {
	return strconv.FormatInt(c.i, 10)
}
