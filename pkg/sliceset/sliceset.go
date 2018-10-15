package sliceset

import "sort"

// After seen the talk at https://www.youtube.com/watch?v=jEG4Qyo_4Bc I wanted
// to check the performance of set-type interface built over a slice, for small
// item counts.

// The idea is to substitute a map[string]struct{} with this if it performs ok
// for the right use case.

// LinearSliceset ...
type LinearSliceset []string

// Insert ...
func (ss *LinearSliceset) Insert(val string) {
	s := *ss
	for _, v := range s {
		if v == val {
			return
		}
	}
	// Not found, just append
	*ss = append(s, val)
}

// IsMember ...
func (ss *LinearSliceset) IsMember(val string) bool {

	for _, v := range *ss {
		if v == val {
			return true
		}
	}
	return false
}

// Delete ...
// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
func (ss *LinearSliceset) Delete(val string) {
	s := *ss
	for i, v := range s {
		if v == val {
			s[i] = s[len(s)-1]
			*ss = s[:len(s)-1]
			return
		}
	}
}

// Snapshot ...
func (ss *LinearSliceset) Snapshot() []string {
	s := *ss
	return append(s[:0:0], s...)
}

// BinarySliceset ...
type BinarySliceset []string

// Insert ...
func (bs *BinarySliceset) Insert(val string) {
	s := *bs
	idx := sort.SearchStrings(s, val)
	if !(idx < len(s) && s[idx] == val) {
		// https://github.com/golang/go/wiki/SliceTricks#insert
		s = append(s, "")
		copy(s[idx+1:], s[idx:])
		s[idx] = val
		*bs = s
	}
}

// IsMember ...
func (bs *BinarySliceset) IsMember(val string) bool {
	s := *bs
	idx := sort.SearchStrings(s, val)
	if idx < len(s) && s[idx] == val {
		return true
	}
	return false
}

// Delete ...
// https://github.com/golang/go/wiki/SliceTricks#delete
func (bs *BinarySliceset) Delete(val string) {
	s := *bs
	idx := sort.SearchStrings(s, val)
	if idx < len(s) && s[idx] == val {
		*bs = append(s[:idx], s[idx+1:]...)
	}
}

// Snapshot ...
func (bs *BinarySliceset) Snapshot() []string {
	s := *bs
	return append(s[:0:0], s...)
}

// HybridSet ...
type HybridSet struct {
	Slice []string
	Set   map[string]int
}

// NewHybridSet ...
func NewHybridSet(hintSize int) *HybridSet {
	ret := &HybridSet{
		Set: make(map[string]int),
	}
	if hintSize != 0 {
		ret.Slice = make([]string, 0, hintSize)
	}
	return ret
}

// Insert ...
func (hs *HybridSet) Insert(val string) {
	if _, ok := hs.Set[val]; ok {
		return
	}
	hs.Slice = append(hs.Slice, val) // Append at the end
	hs.Set[val] = len(hs.Slice)
}

// IsMember ...
func (hs *HybridSet) IsMember(val string) bool {
	_, ok := hs.Set[val]
	return ok
}

// Delete ...
// https://github.com/golang/go/wiki/SliceTricks#delete-without-preserving-order
func (hs *HybridSet) Delete(val string) {
	idx, ok := hs.Set[val]
	if !ok {
		return
	}
	// If only it was this easy... map indexes would move with this method
	// so this is wrong :D
	// Ideas:
	// - Waste more memory with a slice of empty slice positions, so the
	// slice only grows if there hasn't been deletes, but still it would
	// never shrink with this by itself. We could shrink it if the empties
	// slice grows over a max empty size.
	delete(hs.Set, val)
	hs.Slice[idx] = hs.Slice[len(hs.Slice)-1]
	hs.Slice = hs.Slice[:len(hs.Slice)-1]
}

// Snapshot ...
func (hs *HybridSet) Snapshot() []string {
	return append(hs.Slice[:0:0], hs.Slice...)
}
