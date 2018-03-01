package main

import "testing"

var base = []string{
	"benchmarking",
	"is",
	"the",
	"practice",
	"of",
	"measuring",
	"the",
	"performance",
	"of",
	"a",
	"program",
	"on",
	"a",
	"fixed",
	"workload",
}

func BenchmarkWithAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateWithAssign(base...)
	}
}

func BenchmarkWithJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcatenateWithJoin(base...)
	}
}
