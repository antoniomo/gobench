package main

import (
	"testing"
)

const (
	testsize = 1000000
)

var (
	cpyA []int
	cpyB []int
	cpyC []int
	cpyD []int
)

// A, B, C methods from https://github.com/golang/go/wiki/SliceTricks#copy

func BenchmarkSliceCopyA(b *testing.B) {
	s := make([]int, testsize)
	for i := 0; i < testsize; i++ {
		s[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpyA = make([]int, 0, len(s))
		copy(cpyA, s)
	}
}

func BenchmarkSliceCopyB(b *testing.B) {
	s := make([]int, testsize)
	for i := 0; i < testsize; i++ {
		s[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpyB = append([]int(nil), s...)
		cpyB = []int{}
	}
}

func BenchmarkSliceCopyC(b *testing.B) {
	s := make([]int, testsize)
	for i := 0; i < testsize; i++ {
		s[i] = i
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cpyC = append(s[:0:0], s...)
		cpyC = []int{}
	}
}
