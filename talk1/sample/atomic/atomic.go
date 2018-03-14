package main

import (
	"errors"
	"sync/atomic"
)

type counter struct {
	value int64
}

func (c *counter) Add(v int64) {
	if v < 0 {
		panic(errors.New("counter cannot decrease"))
	}
	atomic.AddInt64(&c.value, v)
}

func (c *counter) Write() int64 {
	return atomic.LoadInt64(&c.value)
}
