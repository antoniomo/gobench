package cmap

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"strconv"
	"sync"
	"testing"
)

const maxrangebuf = 10

func sliceranger(l *sync.RWMutex, s []string) <-chan string {
	ch := make(chan string)

	go func() {
		l.RLock()
		defer l.RUnlock()
		for _, str := range s {
			l.RUnlock()
			ch <- str
			l.RLock()
		}
		close(ch)
	}()

	return ch
}

func slicerangerbuf(l *sync.RWMutex, s []string) <-chan string {
	ch := make(chan string, maxrangebuf)

	go func() {
		l.RLock()
		defer l.RUnlock()
		for _, str := range s {
			l.RUnlock()
			ch <- str
			l.RLock()
		}
		close(ch)
	}()

	return ch
}

func BenchmarkBaseline100(b *testing.B) {

	// Setup
	var (
		a []string
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		for _, s := range a {
			_ = s
		}
	}
}

func BenchmarkLocked100(b *testing.B) {

	// Setup
	var (
		a []string
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		for _, s := range a {
			_ = s
		}
		l.RUnlock()
	}
}

func BenchmarkLockDance100(b *testing.B) {

	// Setup
	var (
		a []string
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		for _, s := range a {
			l.RUnlock()
			_ = s
			l.RLock()
		}
		l.RUnlock()
	}
}

func BenchmarkSnapshot100(b *testing.B) {

	// Setup
	var (
		a []string
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		sn := append(a[:0:0], a...)
		l.RUnlock()
		for _, s := range sn {
			_ = s
		}
	}
}

func BenchmarkJsonSnapshot100(b *testing.B) {

	// Setup
	var (
		a []string
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		b, _ := json.Marshal(a)
		l.RUnlock()
		var sn []string
		json.Unmarshal(b, &sn)
		for _, s := range sn {
			_ = s
		}
	}
}

func BenchmarkGobSnapshot100(b *testing.B) {

	// Setup
	var (
		a []string
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		l.RLock()
		enc.Encode(a)
		l.RUnlock()
		r := bytes.NewReader(buf.Bytes())
		dec := gob.NewDecoder(r)
		var sn []string
		dec.Decode(sn)
		for _, s := range sn {
			_ = s
		}
	}
}

func BenchmarkSliceranger100(b *testing.B) {

	// Setup
	var (
		a []string
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		for s := range sliceranger(&l, a) {
			_ = s
		}
	}
}

func BenchmarkSlicerangerBuf100(b *testing.B) {

	// Setup
	var (
		a []string
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		a = append(a, strconv.Itoa(i))
	}

	for i := 0; i < b.N; i++ {
		for s := range slicerangerbuf(&l, a) {
			_ = s
		}
	}
}

func mapranger(l *sync.RWMutex, m map[string]string) <-chan string {
	ch := make(chan string)

	go func() {
		l.RLock()
		defer l.RUnlock()
		for _, str := range m {
			l.RUnlock()
			ch <- str
			l.RLock()
		}
		close(ch)
	}()

	return ch
}

func maprangerbuf(l *sync.RWMutex, m map[string]string) <-chan string {
	ch := make(chan string, maxrangebuf)

	go func() {
		l.RLock()
		defer l.RUnlock()
		for _, str := range m {
			l.RUnlock()
			ch <- str
			l.RLock()
		}
		close(ch)
	}()

	return ch
}

func BenchmarkBaselineMap100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		for _, s := range m {
			_ = s
		}
	}
}

func BenchmarkLockedMap100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		for _, s := range m {
			_ = s
		}
		l.RUnlock()
	}
}

func BenchmarkLockDanceMap100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		for _, s := range m {
			l.RUnlock()
			_ = s
			l.RLock()
		}
		l.RUnlock()
	}
}

func BenchmarkSnapshotMap100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		sn := make(map[string]string, len(m))
		for k, v := range m {
			sn[k] = v
		}
		l.RUnlock()
		for _, s := range sn {
			_ = s
		}
	}
}

func BenchmarkJsonSnapshotMap100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		l.RLock()
		b, _ := json.Marshal(m)
		l.RUnlock()
		var sn map[string]string
		json.Unmarshal(b, &sn)
		for _, s := range sn {
			_ = s
		}
	}
}

func BenchmarkGobSnapshotMap100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := gob.NewEncoder(&buf)
		l.RLock()
		enc.Encode(m)
		l.RUnlock()
		r := bytes.NewReader(buf.Bytes())
		dec := gob.NewDecoder(r)
		var sn []string
		dec.Decode(sn)
		for _, s := range sn {
			_ = s
		}
	}
}

func BenchmarkMapranger100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		for s := range mapranger(&l, m) {
			_ = s
		}
	}
}

func BenchmarkMaprangerBuf100(b *testing.B) {

	// Setup
	var (
		m = make(map[string]string)
		l sync.RWMutex
	)
	for i := 0; i < 100; i++ {
		s := strconv.Itoa(i)
		m[s] = s
	}

	for i := 0; i < b.N; i++ {
		for s := range maprangerbuf(&l, m) {
			_ = s
		}
	}
}
