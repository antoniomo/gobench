package testslice

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/antoniomo/gobench/pkg/sliceset"
)

const (
	testsize = 10
)

var (
	k, v    string
	ok      bool
	testset []int
)

func init() {
	testset = rand.Perm(testsize)
}

func BenchmarkMapInsertNew(b *testing.B) {
	m := make(map[string]struct{})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			m[strconv.Itoa(testset[j])] = struct{}{}
		}
	}
}

func BenchmarkSliceSetInsertNew(b *testing.B) {
	m := sliceset.LinearSliceset{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			m.Insert(strconv.Itoa(testset[j]))
		}
	}
}

func BenchmarkBinarySliceSetInsertNew(b *testing.B) {
	m := sliceset.BinarySliceset{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			m.Insert(strconv.Itoa(testset[j]))
		}
	}
}

func BenchmarkHybridSliceSetInsertNew(b *testing.B) {
	m := sliceset.NewHybridSet(0)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			m.Insert(strconv.Itoa(testset[j]))
		}
	}
}

func BenchmarkHybridSliceSetHintInsertNew(b *testing.B) {
	m := sliceset.NewHybridSet(testsize)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			m.Insert(strconv.Itoa(testset[j]))
		}
	}
}

func BenchmarkMapGet(b *testing.B) {
	m := make(map[string]struct{})

	for i := 0; i < testsize; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			_, ok = m[strconv.Itoa(testset[j])]
		}
	}
}

func BenchmarkSliceSetGet(b *testing.B) {
	m := sliceset.LinearSliceset{}

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			ok = m.IsMember(strconv.Itoa(testset[j]))
		}
	}
}

func BenchmarkBinarySliceSetGet(b *testing.B) {
	m := sliceset.BinarySliceset{}

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			ok = m.IsMember(strconv.Itoa(testset[j]))
		}
	}
}

func BenchmarkHybridSliceSetGet(b *testing.B) {
	m := sliceset.NewHybridSet(0)

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testsize; j++ {
			ok = m.IsMember(strconv.Itoa(testset[j]))
		}
	}
}

// func BenchmarkMapDelete(b *testing.B) {
// 	m := make(map[string]struct{})

// 	for i := 0; i < testsize; i++ {
// 		m[strconv.Itoa(i)] = struct{}{}
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		delete(m, strconv.Itoa(i))
// 	}
// }

// func BenchmarkSliceSetDelete(b *testing.B) {
// 	m := sliceset.LinearSliceset{}

// 	for i := 0; i < testsize; i++ {
// 		m.Insert(strconv.Itoa(i))
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		m.Delete(strconv.Itoa(i))
// 	}
// }

// func BenchmarkBinarySliceSetDelete(b *testing.B) {
// 	m := sliceset.BinarySliceset{}

// 	for i := 0; i < testsize; i++ {
// 		m.Insert(strconv.Itoa(i))
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		m.Delete(strconv.Itoa(i))
// 	}
// }

// func BenchmarkHybridSliceSetDelete(b *testing.B) {
// 	m := sliceset.NewHybridSet(0)

// 	for i := 0; i < testsize; i++ {
// 		m.Insert(strconv.Itoa(i))
// 	}

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		m.Delete(strconv.Itoa(i))
// 	}
// }

func BenchmarkMapRange(b *testing.B) {
	m := make(map[string]struct{})

	for i := 0; i < testsize; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for key := range m {
			k = key
		}
	}
}

func BenchmarkSliceSetRange(b *testing.B) {
	m := sliceset.LinearSliceset{}

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, key := range m {
			k = key
		}
	}
}

func BenchmarkBinarySliceSetRange(b *testing.B) {
	m := sliceset.BinarySliceset{}

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, key := range m {
			k = key
		}
	}
}

func BenchmarkHybridSliceSetRange(b *testing.B) {
	m := sliceset.NewHybridSet(0)

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, key := range m.Slice {
			k = key
		}
	}
}

func BenchmarkMapSnapshotRange(b *testing.B) {
	m := make(map[string]struct{})

	b.ResetTimer()
	for i := 0; i < testsize; i++ {
		m[strconv.Itoa(i)] = struct{}{}
	}

	for i := 0; i < b.N; i++ {
		snapshot := make([]string, 0, len(m))
		for k := range m {
			snapshot = append(snapshot, k)
		}
		for _, key := range snapshot {
			k = key
		}
	}
}

func BenchmarkLinearSliceSetSnapshotRange(b *testing.B) {
	m := sliceset.LinearSliceset{}

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		snapshot := m.Snapshot()
		for _, key := range snapshot {
			k = key
		}
	}
}
func BenchmarkBinarySliceSetSnapshotRange(b *testing.B) {
	m := sliceset.BinarySliceset{}

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		snapshot := m.Snapshot()
		for _, key := range snapshot {
			k = key
		}
	}
}
func BenchmarkHybridSnapshotRange(b *testing.B) {
	m := sliceset.NewHybridSet(0)

	for i := 0; i < testsize; i++ {
		m.Insert(strconv.Itoa(i))
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		snapshot := m.Snapshot()
		for _, key := range snapshot {
			k = key
		}
	}
}
