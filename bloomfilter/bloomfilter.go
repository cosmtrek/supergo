package bloomfilter

import (
	"encoding/binary"
	"hash/fnv"
	"math"

	"github.com/cosmtrek/supergo/bitset"
)

// BloomFilter ...
type BloomFilter struct {
	m   uint // holds number
	set *bitset.BitSet
	p   float64 // error tolerant rate
}

// NewBloomFilter ...
func NewBloomFilter(m uint, p float64) *BloomFilter {
	nBits := uint(math.Ceil(float64(m) * math.Log(1/p) / math.Pow(math.Log(2), 2)))
	return &BloomFilter{
		m:   m,
		set: bitset.NewBitSet(nBits),
	}
}

// Add ...
func (b *BloomFilter) Add(key []byte) {
	h := fnv.New64()
	h.Write(key)
	s := h.Sum(nil)
	for i := 0; i < h.Size()/4; i++ {
		k := uint(binary.BigEndian.Uint32(s[i*4:])) % b.m
		b.set.Set(k)
	}
}

// HasKey ...
func (b *BloomFilter) HasKey(key []byte) bool {
	h := fnv.New64()
	h.Write(key)
	s := h.Sum(nil)
	for i := 0; i < h.Size()/4; i++ {
		k := uint(binary.BigEndian.Uint32(s[i*4:])) % b.m
		if !b.set.Has(k) {
			return false
		}
	}
	return true
}
