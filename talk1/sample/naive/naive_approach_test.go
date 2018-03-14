package main

import "testing"

func BenchmarkNaiveRoutineCounter100(b *testing.B) {
	c := counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 100)
}

func BenchmarkNaiveRoutineCounter10(b *testing.B) {
	c := counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 10)
}

func BenchmarkNaiveRoutineCounter1(b *testing.B) {
	c := counter{}
	benchmarkAddAndWriteWithRoutine(b, c, 1)
}
