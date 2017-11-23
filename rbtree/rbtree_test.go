package rbtree

import (
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := NewTree()
	if tree.root != nil {
		t.Error("Failed to initialize tree!")
	}
}

func TestInsert(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	if tree.root.key != 10 {
		t.Error("Failed to insert node 10!")
	}
	if tree.root.color != BLACK {
		t.Error("Node 10 wrong color!")
	}

	tree.Insert(8)
	if tree.root.left.key != 8 {
		t.Error("Failed to insert node 8!")
	}
	if tree.root.left.color != RED {
		t.Error("Node 8 wrong color!")
	}

	tree.Insert(12)
	if tree.root.right.key != 12 {
		t.Error("Failed to insert node 12!")
	}
	if tree.root.right.color != RED {
		t.Error("Node 12 wrong color!")
	}
}

func exist(t *Tree, key int) bool {
	if _, ok := t.Search(key); !ok {
		return false
	}

	return true
}

func TestSearch(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(0)
	tree.Insert(1)
	tree.Insert(13)
	tree.Insert(12)
	tree.Insert(20)
	if exist(tree, 100) {
		t.Errorf("No node %d", 100)
	}
	if !exist(tree, 0) {
		t.Errorf("Failed to find node %d", 0)
	}
	if !exist(tree, 10) {
		t.Errorf("Failed to find node %d", 10)
	}
	if !exist(tree, 20) {
		t.Errorf("Failed to find node %d", 20)
	}
}

func TestDelete(t *testing.T) {
	tree := NewTree()
	tree.Insert(10)
	tree.Insert(40)
	tree.Insert(30)
	tree.Insert(60)
	tree.Insert(90)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(50)
	if exist(tree, 100) {
		t.Errorf("No node %d", 100)
	}

	if !exist(tree, 40) {
		t.Errorf("Failed to find node %d", 40)
	}
	tree.Delete(40)
	if exist(tree, 40) {
		t.Errorf("Failed to delete node %d", 40)
	}

	if !exist(tree, 30) {
		t.Errorf("Failed to find node %d", 30)
	}
	tree.Delete(30)
	if exist(tree, 30) {
		t.Errorf("Failed to delete node %d", 30)
	}

	if !exist(tree, 20) {
		t.Errorf("Failed to find node %d", 20)
	}
	tree.Delete(20)
	if exist(tree, 20) {
		t.Errorf("Failed to delete node %d", 20)
	}
}
