package leakybuffer

import (
	"log"
)

// LeakyBuffer ...
type LeakyBuffer struct {
	bufSize  int         // each buffer size
	freeList chan []byte // n buffers
}

// NewLeakyBuffer ...
func NewLeakyBuffer(bufSize, nBuf int) *LeakyBuffer {
	return &LeakyBuffer{
		bufSize:  bufSize,
		freeList: make(chan []byte, nBuf),
	}
}

// Get ...
func (b *LeakyBuffer) Get() []byte {
	select {
	case nb := <-b.freeList:
		return nb
	default:
		return make([]byte, b.bufSize)
	}
}

// Put ...
func (b *LeakyBuffer) Put(n []byte) {
	if len(n) != b.bufSize {
		log.Fatal("Invalid buffer size for LeakyBuffer!")
	}

	select {
	case b.freeList <- n:
	}
}
