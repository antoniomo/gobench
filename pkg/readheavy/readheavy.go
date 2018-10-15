package readheavy

import (
	"sync"
	"sync/atomic"
)

// Adapted from https://golang.org/pkg/sync/atomic/#example_Value_readMostly

type innerMap map[string]string

// Map ...
type Map struct {
	av atomic.Value
	mu sync.Mutex // used only by writers
	m  innerMap
}

// New ...
func New() *Map {

	ret := &Map{
		m: make(innerMap),
	}
	ret.av.Store(ret.m)

	return ret
}

// Get ...
func (m *Map) Get(key string) (string, bool) {
	m1 := m.av.Load().(innerMap)
	value, ok := m1[key]
	return value, ok
}

// Set ...
func (m *Map) Set(key, value string) {
	m.mu.Lock()                  // synchronize with other potential writers
	m1 := m.av.Load().(innerMap) // load current value of the data structure
	m2 := make(innerMap)         // create a new value
	for k, v := range m1 {
		m2[k] = v // copy all data from the current object to the new one
	}
	m2[key] = value // do the update that we need
	m.av.Store(m2)  // atomically replace the current object with the new one
	m.mu.Unlock()
	// At this point all new readers start working with the new version.
	// The old version will be garbage collected once the existing readers
	// (if any) are done with it.
}

// Delete ...
func (m *Map) Delete(key string) {
	m.mu.Lock()                  // synchronize with other potential writers
	m1 := m.av.Load().(innerMap) // load current value of the data structure
	m2 := make(innerMap)         // create a new value
	delete(m1, key)              // do the update that we need
	for k, v := range m1 {
		m2[k] = v // copy all data from the current object to the new one
	}
	m.av.Store(m2) // atomically replace the current object with the new one
	m.mu.Unlock()
	// At this point all new readers start working with the new version.
	// The old version will be garbage collected once the existing readers
	// (if any) are done with it.
}

// Tuple ...
type Tuple struct {
	Key   string
	Value string
}

// Iter to use with a range loop
func (m *Map) Iter() <-chan Tuple {

	snap := m.av.Load().(innerMap)
	// Fully buffered, trade memory for speed baby!
	ret := make(chan Tuple, len(snap))
	go func() {
		for k, v := range snap {
			ret <- Tuple{k, v}
		}
		close(ret)
	}()
	return ret
}

// Snapshot ...
func (m *Map) Snapshot() map[string]string {

	return m.av.Load().(innerMap)
}

// Extend ...
func (m *Map) Extend(e map[string]string) {
	m.mu.Lock()                  // synchronize with other potential writers
	m1 := m.av.Load().(innerMap) // load current value of the data structure
	m2 := make(innerMap)         // create a new value
	for k, v := range m1 {
		m2[k] = v // copy all data from the current object to the new one
	}
	for k, v := range e {
		m2[k] = v // copy all data from the e map too
	}
	m.av.Store(m2) // atomically replace the current object with the new one
	m.mu.Unlock()
	// At this point all new readers start working with the new version.
	// The old version will be garbage collected once the existing readers
	// (if any) are done with it.
}

// ExtendSlice ...
func (m *Map) ExtendSlice(e []Tuple) {
	m.mu.Lock()                  // synchronize with other potential writers
	m1 := m.av.Load().(innerMap) // load current value of the data structure
	m2 := make(innerMap)         // create a new value
	for k, v := range m1 {
		m2[k] = v // copy all data from the current object to the new one
	}
	for _, kv := range e {
		m2[kv.Key] = kv.Value // copy all data from the e slice too
	}
	m.av.Store(m2) // atomically replace the current object with the new one
	m.mu.Unlock()
	// At this point all new readers start working with the new version.
	// The old version will be garbage collected once the existing readers
	// (if any) are done with it.
}
