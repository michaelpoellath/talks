package main

import "testing"

func BenchmarkNaiveCounter(b *testing.B) {
	c := counter{}
	benchmarkAddAndWrite(b, c)
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
