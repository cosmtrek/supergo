package bloomfilter

import (
	"strconv"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	b := NewBloomFilter(100000000, 0.0001)
	key := []byte("golang")
	b.Add(key)
	if !b.HasKey(key) {
		t.Error("Failed to add key to the bloomfilter")
	}
	noKey := []byte("bloomfilter")
	if b.HasKey(noKey) {
		t.Error("May not be in the set")
	}
}

func BenchmarkBloomFilter(b *testing.B) {
	f := NewBloomFilter(100000000, 0.0001)
	for i := 0; i < 100000000; i++ {
		f.Add([]byte(strconv.Itoa(i)))
	}
	for i := 0; i < 100000000; i++ {
		if !f.HasKey([]byte(strconv.Itoa(i))) {
			b.Error("May not be in the set")
		}
	}
}
