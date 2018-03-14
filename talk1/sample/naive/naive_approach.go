package main

import "errors"

type counter struct {
	value int
}

func (c *counter) Add(v int) {
	if v < 0 {
		panic(errors.New("counter cannot decrease"))
	}
	c.value += v
}

func (c counter) Write() int {
	return c.value
}
