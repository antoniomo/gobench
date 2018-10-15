package testslice

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/antoniomo/gobench/pkg/slicemap"
)

const (
	testsize = 500
)

var (
	k, v        string
	ok          bool
	testindexes []int
)

func BenchmarkMapInsertNew(b *testing.B) {
	m := make(map[string]string)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(i)] = "asdfasdf"
	}
}

func BenchmarkSliceMapInsertNew(b *testing.B) {
	m := make(slicemap.LinearSlicemap, 0, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
}

func BenchmarkBinarySliceMapInsertNew(b *testing.B) {
	m := make(slicemap.BinarySlicemap, 0, b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
}

func BenchmarkMapInsertNewRandomCase(b *testing.B) {
	m := make(map[string]string)
	testindexes := rand.Perm(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[strconv.Itoa(testindexes[i])] = "asdfasdf"
	}
}

func BenchmarkSliceMapInsertNewRandomCase(b *testing.B) {
	m := make(slicemap.LinearSlicemap, 0, b.N)
	testindexes := rand.Perm(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(testindexes[i]), "asdfasdf")
	}
}

func BenchmarkBinarySliceMapInsertNewRandomCase(b *testing.B) {
	m := make(slicemap.BinarySlicemap, 0, b.N)
	testindexes := rand.Perm(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(testindexes[i]), "asdfasdf")
	}
}

func BenchmarkMapGet(b *testing.B) {
	m := make(map[string]string)

	for i := 0; i < testsize; i++ {
		m[strconv.Itoa(i)] = "asdfasdf"
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, ok = m[strconv.Itoa(i)]
	}
}

func BenchmarkSliceMapGet(b *testing.B) {
	m := slicemap.LinearSlicemap{}

	for i := 0; i < testsize; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, ok = m.Get(strconv.Itoa(i))
	}
}

func BenchmarkBinarySliceMapGet(b *testing.B) {
	m := slicemap.BinarySlicemap{}

	for i := 0; i < testsize; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, ok = m.Get(strconv.Itoa(i))
	}
}

func BenchmarkMapRange(b *testing.B) {
	m := make(map[string]string)

	for i := 0; i < testsize; i++ {
		m[strconv.Itoa(i)] = "asdfasdf"
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for key, val := range m {
			k, v = key, val
		}
	}
}

func BenchmarkSliceMapRange(b *testing.B) {
	m := slicemap.LinearSlicemap{}

	for i := 0; i < testsize; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range m {
			k, v = kv.Key, kv.Value
		}
	}
}

func BenchmarkBinarySliceMapRange(b *testing.B) {
	m := slicemap.BinarySlicemap{}

	for i := 0; i < testsize; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range m {
			k, v = kv.Key, kv.Value
		}
	}
}
