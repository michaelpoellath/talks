package main

import "testing"

func BenchmarkMutexCounter100(b *testing.B) {
	c := &counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 100)
}

func BenchmarkMutexCounter10(b *testing.B) {
	c := &counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 10)
}

func BenchmarkMutexCounter1(b *testing.B) {
	c := &counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 1)
}
