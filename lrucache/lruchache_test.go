package lrucache

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestNewLRUCache(t *testing.T) {
	maxSize := 10
	cache := newLRUCache(maxSize)
	if cache.MaxSize() != maxSize {
		t.Errorf("Failed to initialize LRUCache, expected max size: %d, actual: %d", maxSize, cache.MaxSize())
	}
	if cache.Size() != 0 {
		t.Errorf("Failed to initialize LRUCache, expected size: %d, actual: %d", 0, cache.Size())
	}
}

func TestNewLRUCacheWithInvalidParams(t *testing.T) {
	maxSize := -1
	if c := newLRUCache(maxSize); c != nil {
		t.Errorf("Should not initialize cache for invalid max size: %d", maxSize)
	}
}

func TestLRUCacheSetAndGet(t *testing.T) {
	cache := newLRUCache(10)
	k, v := "key", "value"
	val, ok := cache.Set(k, v).Get(k)
	if !ok || val != v {
		t.Errorf("Failed to Get key: %s, value expected: %s, actual: %s", k, v, val)
	}
}

func TestLRUCacheExpireOldEntry(t *testing.T) {
	cache := newLRUCache(3)
	k1, v1 := "key1", "value1"
	k2, v2 := "key2", "value2"
	k3, v3 := "key3", "value3"
	k4, v4 := "key4", "value4"
	cache.Set(k1, v1)
	cache.Set(k2, v2)
	cache.Set(k3, v3)
	cache.Set(k4, v4)
	if _, ok := cache.Get(k1); ok {
		t.Errorf("Failed to expire the oldest entry, key: %s, value: %s", k1, v1)
	}
}

func BenchmarkLRUCache(b *testing.B) {
	n := 1000000
	cache := newLRUCache(10000)

	for i := 0; i < n; i++ {
		r1 := rand.Int() * n
		r2 := rand.Int() * n
		k, v := fmt.Sprintf("Key%d", r1), fmt.Sprintf("Value%d", r1)
		cache.Set(k, v)
		cache.Get(fmt.Sprintf("Key%d", r2))
	}
}
