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

func BenchmarkChannelCounter100(b *testing.B) {
	c := &counter{
		in:  make(chan int),
		out: make(chan int),
	}
	benchmarkAddAndWriteWithRoutine(b, c, 100)

}

func BenchmarkChannelCounter10(b *testing.B) {
	in := make(chan int)
	out := make(chan int)
	c := &counter{
		in:  in,
		out: out,
	}
	benchmarkAddAndWriteWithRoutine(b, c, 10)
}

func BenchmarkChannelCounter1(b *testing.B) {
	c := &counter{
		in:  make(chan int),
		out: make(chan int),
	}
	benchmarkAddAndWriteWithRoutine(b, c, 1)
}
