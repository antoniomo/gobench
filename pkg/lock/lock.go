package lock

import "sync"

// Map ...
type Map struct {
	mu sync.Mutex
	m  map[string]string
}

// New ...
func New() *Map {
	return &Map{m: make(map[string]string)}
}

// Get ...
func (m *Map) Get(key string) (string, bool) {
	m.mu.Lock()
	value, ok := m.m[key]
	m.mu.Unlock()
	return value, ok
}

// Set ...
func (m *Map) Set(key, value string) {
	m.mu.Lock()
	m.m[key] = value
	m.mu.Unlock()
}

// Delete ...
func (m *Map) Delete(key string) {
	m.mu.Lock()
	delete(m.m, key)
	m.mu.Unlock()
}

// Tuple ...
type Tuple struct {
	Key   string
	Value string
}

// Iter to use with a range loop
func (m *Map) Iter() <-chan Tuple {

	snap := m.Snapshot()
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

// IterSlice to use with a range loop (with slice snapshot)
func (m *Map) IterSlice() <-chan Tuple {

	snap := m.SliceSnapshot()
	// Fully buffered, trade memory for speed baby!
	ret := make(chan Tuple, len(snap))
	go func() {
		for _, v := range snap {
			ret <- v
		}
		close(ret)
	}()
	return ret
}

// Snapshot ...
func (m *Map) Snapshot() map[string]string {

	ret := make(map[string]string)
	m.mu.Lock()
	for k, v := range m.m {
		ret[k] = v
	}
	m.mu.Unlock()
	return ret
}

// SliceSnapshot ...
func (m *Map) SliceSnapshot() []Tuple {
	i := 0
	m.mu.Lock()
	ret := make([]Tuple, len(m.m))
	for k, v := range m.m {
		ret[i] = Tuple{k, v}
		i++
	}
	m.mu.Unlock()
	return ret
}
