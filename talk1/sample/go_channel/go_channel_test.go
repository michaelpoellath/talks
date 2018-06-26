package main

import "testing"

func BenchmarkChannelCounter100(b *testing.B) {
	c := &counter{
		in:  make(chan int),
		out: make(chan int),
	}
	benchmarkAddAndWriteWithRoutine(b, c, 100)
}

func BenchmarkChannelCounter10(b *testing.B) {
	c := &counter{
		in:  make(chan int),
		out: make(chan int),
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
