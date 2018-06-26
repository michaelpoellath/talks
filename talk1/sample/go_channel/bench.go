package main

import (
	"fmt"
	"sync"
	"testing"
)

//Benchmark with go routines
func benchmarkAddAndWriteWithRoutine(b *testing.B, c *counter, concurrency int) {
	b.StopTimer()
	var start, end sync.WaitGroup
	start.Add(1)
	end.Add(concurrency)
	n := b.N / concurrency

	go c.loop()
	for i := 0; i < concurrency; i++ {
		go func() {
			start.Wait()
			for i := 0; i < n; i++ {
				if i%1000 == 0 {
					t := c.Write()
					t = t + 1
				}
				c.Add(1)
			}
			end.Done()
		}()
	}
	b.StartTimer()
	start.Done()
	end.Wait()
	fmt.Printf("Iterations: %d  - fetched from counter: %d\n", b.N, c.Write())

}
