package main

import "testing"

func BenchmarkHealthWithStatsHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateStats()
	}
}
