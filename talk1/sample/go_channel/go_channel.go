package main

type counter struct {
	in  chan int //May be buffered
	out chan int //Must be synchronous.
}

func (c *counter) Add(v int) {
	c.in <- v
}

func (c *counter) Write() int {
	return <-c.out
}

func (c *counter) loop() {
	var value int
	for {
		select {
		case v := <-c.in:
			value += v
		case c.out <- value:
			//Do Nothing
		}
	}
}
