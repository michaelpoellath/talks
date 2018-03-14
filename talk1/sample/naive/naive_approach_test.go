package main

import (
	"fmt"
	"sync"
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

func BenchmarkNaiveCounter(b *testing.B) {
	c := counter{}
	benchmarkAddAndWrite(b, c)
}

//Benchmark with go routines
func benchmarkAddAndWriteWithRoutine(b *testing.B, c counter, concurrency int) {
	b.StopTimer()
	var start, end sync.WaitGroup
	start.Add(1)
	end.Add(concurrency)
	n := b.N / concurrency

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

func BenchmarkNaiveCounter100(b *testing.B) {
	c := counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 100)
}

func BenchmarkNaiveCounter10(b *testing.B) {
	c := counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 10)
}

func BenchmarkNaiveCounter1(b *testing.B) {
	c := counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 1)
}
