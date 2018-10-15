package main

import (
	"strconv"
	"sync"
	"testing"

	"github.com/antoniomo/gobench/pkg/lock"
	"github.com/antoniomo/gobench/pkg/readheavy"
	"github.com/antoniomo/gobench/pkg/rwlock"
)

const (
	testelements = 1000
)

var (
	K string
	V string
)

func BenchmarkLockInsert(b *testing.B) {
	m := lock.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
}

func BenchmarkRWLockInsert(b *testing.B) {
	m := rwlock.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
}

func BenchmarkSyncMapInsert(b *testing.B) {
	m := sync.Map{}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Store(strconv.Itoa(i), "asdfasdf")
	}
}

func BenchmarkReadHeavyInsert(b *testing.B) {
	m := readheavy.New()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
}

func BenchmarkReadHeavyExtend(b *testing.B) {
	m := readheavy.New()
	mm := make(map[string]string)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testelements; j++ {
			mm[strconv.Itoa(i)] = "asdfasdf"
		}
		m.Extend(mm)
	}
}

func BenchmarkReadHeavyExtendSlice(b *testing.B) {
	m := readheavy.New()
	mm := make([]readheavy.Tuple, testelements)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < testelements; j++ {
			mm[j] = readheavy.Tuple{Key: strconv.Itoa(i), Value: "asdfasdf"}
		}
		m.ExtendSlice(mm)
	}
}

func BenchmarkLockIter(b *testing.B) {
	m := lock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for kv := range m.Iter() {
			K, V = kv.Key, kv.Value
		}
	}
}

func BenchmarkLockIterSlice(b *testing.B) {
	m := lock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for kv := range m.IterSlice() {
			K, V = kv.Key, kv.Value
		}
	}
}

func BenchmarkLockSnapRange(b *testing.B) {
	m := lock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for k, v := range m.Snapshot() {
			K, V = k, v
		}
	}
}

func BenchmarkLockSliceRange(b *testing.B) {
	m := lock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range m.SliceSnapshot() {
			K, V = kv.Key, kv.Value
		}
	}
}

func BenchmarkRWLockIter(b *testing.B) {
	m := rwlock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for kv := range m.Iter() {
			K, V = kv.Key, kv.Value
		}
	}
}

func BenchmarkRWLockIterSlice(b *testing.B) {
	m := rwlock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for kv := range m.IterSlice() {
			K, V = kv.Key, kv.Value
		}
	}
}

func BenchmarkRWLockSnapRange(b *testing.B) {
	m := rwlock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for k, v := range m.Snapshot() {
			K, V = k, v
		}
	}
}

func BenchmarkRWLockSliceRange(b *testing.B) {
	m := rwlock.New()

	for i := 0; i < testelements; i++ {
		m.Set(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, kv := range m.SliceSnapshot() {
			K, V = kv.Key, kv.Value
		}
	}
}

func BenchmarkSyncMapRange(b *testing.B) {
	m := sync.Map{}

	for i := 0; i < testelements; i++ {
		m.Store(strconv.Itoa(i), "asdfasdf")
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m.Range(func(k, v interface{}) bool {
			K, V = k.(string), v.(string)
			return true
		})
	}
}

func BenchmarkReadHeavyIter(b *testing.B) {
	m := readheavy.New()
	mm := make(map[string]string)

	for i := 0; i < testelements; i++ {
		mm[strconv.Itoa(i)] = "asdfasdf"
	}
	m.Extend(mm)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for kv := range m.Iter() {
			K, V = kv.Key, kv.Value
		}
	}
}

func BenchmarkReadHeavySnapRange(b *testing.B) {
	m := readheavy.New()
	mm := make(map[string]string)

	for i := 0; i < testelements; i++ {
		mm[strconv.Itoa(i)] = "asdfasdf"
	}
	m.Extend(mm)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for k, v := range m.Snapshot() {
			K, V = k, v
		}
	}
}
