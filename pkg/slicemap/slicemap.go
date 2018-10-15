package slicemap

import "sort"

// After seen the talk at https://www.youtube.com/watch?v=jEG4Qyo_4Bc I wanted
// to check the performance of map-type interface built over a slice, for small
// item counts.

// Notice that I'm doing a "map" of [string]string, so comparisons aren't cheap
// and normal maps have advantage.

// Tuple ...
type Tuple struct {
	Key   string
	Value string
}

// LinearSlicemap ...
type LinearSlicemap []Tuple

// Set ...
func (sm *LinearSlicemap) Set(k, val string) {
	s := *sm
	for i, v := range s {
		if v.Key == k {
			v.Value = val
			s[i] = v
			return
		}
	}
	// Not found, just append
	*sm = append(s, Tuple{Key: k, Value: val})
}

// Get ...
func (sm *LinearSlicemap) Get(k string) (string, bool) {

	for _, v := range *sm {
		if v.Key == k {
			return v.Value, true
		}
	}
	return "", false
}

// Delete ...
// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
func (sm *LinearSlicemap) Delete(k string) {
	s := *sm
	for i, v := range s {
		if v.Key == k {
			s[i] = s[len(s)-1]
			*sm = s[:len(s)-1]
			return
		}
	}
}

// BinarySlicemap ...
type BinarySlicemap []Tuple

// Set ...
func (bs *BinarySlicemap) Set(k, val string) {
	s := *bs
	idx := sort.Search(len(s), func(i int) bool {
		return s[i].Key >= k
	})
	if idx < len(s) && s[idx].Key == k {
		s[idx].Value = val
	} else {
		// https://github.com/golang/go/wiki/SliceTricks#insert
		s = append(s, Tuple{})
		copy(s[idx+1:], s[idx:])
		s[idx] = Tuple{Key: k, Value: val}
	}
	*bs = s
}

// Get ...
func (bs *BinarySlicemap) Get(k string) (string, bool) {
	s := *bs
	idx := sort.Search(len(s), func(i int) bool {
		return s[i].Key >= k
	})
	if idx < len(s) && s[idx].Key == k {
		return s[idx].Value, true
	}
	return "", false
}

// Delete ...
// https://github.com/golang/go/wiki/SliceTricks#delete
func (bs *BinarySlicemap) Delete(k string) {
	s := *bs
	idx := sort.Search(len(s), func(i int) bool {
		return s[i].Key >= k
	})
	if idx < len(s) && s[idx].Key == k {
		*bs = append(s[:idx], s[idx+1:]...)
	}
}
