package leakybuffer

import (
	"testing"
)

func TestNewLeakyBuffer(t *testing.T) {
	lb := NewLeakyBuffer(16, 10)
	expectedSize := 16
	actualSize := lb.bufSize
	if expectedSize != actualSize {
		t.Errorf("Expected: %d, actual: %d", expectedSize, actualSize)
	}

	expectedBuffers := 10
	actualBuffers := cap(lb.freeList)
	if expectedBuffers != actualBuffers {
		t.Errorf("Expected: %d, actual: %d", expectedBuffers, actualBuffers)
	}
}

func TestGetBuffer(t *testing.T) {
	lb := NewLeakyBuffer(16, 10)
	b := lb.Get()
	expectedSize := lb.bufSize
	actualSize := len(b)
	if expectedSize != actualSize {
		t.Errorf("Expected: %d, actual: %d", expectedSize, actualSize)
	}
}

func TestPutBuffer(t *testing.T) {
	lb := NewLeakyBuffer(16, 10)
	b := lb.Get()
	lb.Put(b)
	expectedSize := lb.bufSize
	actualSize := len(b)
	if expectedSize != actualSize {
		t.Errorf("Expected: %d, actual: %d", expectedSize, actualSize)
	}
}
