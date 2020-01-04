package models

import (
	"time"
)

type Component struct {
}

type Results interface {
}

func (c *Component) GetTimeNow() string {

	t := time.Now()
	return string(t.Format("2006-01-02 15:04:05.999999"))
}
