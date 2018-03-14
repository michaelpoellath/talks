package main

import "sync"

type counter struct {
	value int
	mut   sync.Mutex
}

func (c *counter) Add(v int) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.value += v
}

func (c *counter) Write() int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.value
}
