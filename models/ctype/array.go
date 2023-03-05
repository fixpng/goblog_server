package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (t *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*t = []string{}
		return nil
	}
	*t = strings.Split(string(v), "\n")
	return nil
}

func (t Array) Value() (driver.Value, error) {
	// 将数字转换为值
	return strings.Join(t, "\n"), nil
}
