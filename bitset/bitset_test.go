package bitset

import (
	"testing"
)

func TestNewBitSet(t *testing.T) {
	s := NewBitSet(64)
	s.Set(10)
	if !s.Has(10) {
		t.Errorf("Expected true, actual false, bitset: %v", s.set)
	}
}

func TestClearBit(t *testing.T) {
	s := NewBitSet(64)
	s.Set(8).Clear(8)
	if s.Has(8) {
		t.Errorf("Failed to clear the bit, bitset: %v", s.set)
	}
}

func TestWordNeed(t *testing.T) {
	if wordNeed(64) != 2 {
		t.Errorf("Wrong words, actual: %d", wordNeed(64))
	}
}

func TestAutoShrink(t *testing.T) {
	s := NewBitSet(64)
	if s.Len() != 64 {
		t.Error("Wrong bitset length")
	}
	s.Set(100)
	if s.Len() != 100 {
		t.Error("Failed to autoshrink")
	}
}

func BenchmarkHas(b *testing.B) {
	s := NewBitSet(10000000000)
	for i := 0; i < int(s.Len()); i = i + 2 {
		s.Set(uint(i))
	}
	for i := 0; i < int(s.Len()); i++ {
		if i%2 == 0 {
			if !s.Has(uint(i)) {
				b.Error("Wrong")
			}
		} else {
			continue
		}
	}
}
