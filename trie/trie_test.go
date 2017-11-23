package trie

import (
	"strconv"
	"testing"
)

func TestTrieAdd(t *testing.T) {
	tree := NewTrie()
	v := "a"
	tree.Add(v)
	if !tree.Found(v) {
		t.Errorf("Failed to found value %v", v)
	}
	if tree.Found("aa") {
		t.Errorf("Actually there are no such value: %v", "aa")
	}
}

func TestTrieDestroy(t *testing.T) {
	tree := NewTrie()
	v := "a"
	tree.Add(v)
	if tree.Empty() {
		t.Error("Failed to add node")
	}
	tree.Destroy()
	if !tree.Empty() {
		t.Error("Failed to destroy the tree")
	}
}

func BenchmarkAdd(b *testing.B) {
	tree := NewTrie()
	for i := 0; i < 100000000; i++ {
		tree.Add(strconv.Itoa(i))
	}
}
