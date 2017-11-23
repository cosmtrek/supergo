package bitset

const (
	size    = uint(64)
	lg2size = uint(6)
)

type bits uint64

// BitSet ...
type BitSet struct {
	set    []bits
	length uint
}

// NewBitSet ...
func NewBitSet(n uint) *BitSet {
	return &BitSet{
		set:    make([]bits, wordNeed(n)),
		length: n,
	}
}

func wordNeed(n uint) uint {
	return ((n + size - 1) >> lg2size) + 1
}

// Len returns size of bit set
func (s *BitSet) Len() uint {
	return s.length
}

func (s *BitSet) autoShrink(i uint) *BitSet {
	if i >= s.length {
		newSet := make([]bits, wordNeed(i))
		copy(newSet, s.set)
		s.set = newSet
		s.length = i
	}
	return s
}

// Set ...
func (s *BitSet) Set(i uint) *BitSet {
	s.autoShrink(i)
	s.set[i>>lg2size] |= 1 << (i & (size - 1))
	return s
}

// Has ...
func (s *BitSet) Has(i uint) bool {
	return s.set[i>>lg2size]&(1<<(i&(size-1))) != 0
}

// Clear ...
func (s *BitSet) Clear(i uint) *BitSet {
	if i >= s.length {
		return s
	}
	s.set[i>>lg2size] &^= 1 << (i & (size - 1))
	return s
}
