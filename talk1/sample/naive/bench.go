package main

import (
	"fmt"
	"testing"
)

//Benchmark without go routines
func benchmarkAddAndWrite(b *testing.B, c counter) {
	for i := 0; i < b.N; i++ {
		if i%1000 == 0 {
			t := c.Write()
			t = t + 1
			continue
		}
		c.Add(1)
	}
	fmt.Printf("Iterations: %d  - fetched from counter: %d\n", b.N, c.Write())
}
