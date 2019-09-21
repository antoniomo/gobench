package main

import (
	"testing"
)

const (
	testsize = 1000
)

var (
	s []int
)

func BenchmarkForIdx(b *testing.B) {
	m := make(map[int]struct{})

	for i := 0; i < testsize; i++ {
		m[i] = struct{}{}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = make([]int, testsize)
		x := 0
		for k := range m {
			s[x] = k
			x++
		}
	}
}

func BenchmarkForAppend(b *testing.B) {
	m := make(map[int]struct{})

	for i := 0; i < testsize; i++ {
		m[i] = struct{}{}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = make([]int, 0, testsize)
		for k := range m {
			s = append(s, k)
		}
	}
}

func BenchmarkForIdxReversed(b *testing.B) {
	m := make(map[int]struct{})

	for i := 0; i < testsize; i++ {
		m[i] = struct{}{}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = make([]int, testsize)
		x := testsize - 1
		for k := range m {
			s[x] = k
		}
	}
}
