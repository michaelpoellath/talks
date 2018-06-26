package main

import (
	"errors"
	"sync"
)

type counter struct {
	value int
	mut   sync.Mutex
}

func (c *counter) Add(v int) {
	c.mut.Lock()
	defer c.mut.Unlock()
	if v < 0 {
		panic(errors.New("counter cannot decrease"))
	}
	c.value += v
}

func (c *counter) Write() int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.value
}
