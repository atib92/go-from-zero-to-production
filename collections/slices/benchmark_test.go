package main

import "testing"

func BenchmarkAppendPreallocated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 1000)

		for j := 0; j < 1000; j++ {
			s = append(s, j)
		}
	}
}

func BenchmarkAppendWithoutPreallocation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var s []int

		for j := 0; j < 1000; j++ {
			s = append(s, j)
		}
	}
}