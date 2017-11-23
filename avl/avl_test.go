package avl

import (
	"math/rand"
	"testing"
)

func TestNewAVL(t *testing.T) {
	tree := NewTree(1)
	if tree.root.value != 1 {
		t.Errorf("Failed to initialize AVL tree!")
	}
}

func TestInsertNode(t *testing.T) {
	tree := NewTree(0)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(-1)
	r := tree.root
	if r.height != 2 {
		t.Errorf("Failed! Tree: %v, left: %v, right: %v", r, r.left, r.right)
	}
}

// LL case
func TestRightRotate(t *testing.T) {
	tree := NewTree(8)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(2)
	tree.Insert(6)
	if tree.root.height != 2 {
		t.Error("Failed to insert node!")
	}
	if tree.root.bfactor() != 1 {
		t.Error("Wrong bfactor!")
	}
	tree.Insert(1)
	if tree.root.bfactor() != 0 {
		t.Errorf("Failed to LL! root: %v", tree.root)
	}
}

// RR case
func TestLeftRotate(t *testing.T) {
	tree := NewTree(8)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(10)
	tree.Insert(14)
	if tree.root.height != 2 {
		t.Error("Failed to insert node!")
	}
	if tree.root.bfactor() != -1 {
		t.Error("Wrong bfactor!")
	}
	tree.Insert(13)
	if tree.root.bfactor() != 0 {
		t.Errorf("Failed to LL! root: %v", tree.root)
	}
}

// LR case
func TestLeftRightRotate(t *testing.T) {
	tree := NewTree(8)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(2)
	tree.Insert(6)
	if tree.root.height != 2 {
		t.Error("Failed to insert node!")
	}
	if tree.root.bfactor() != 1 {
		t.Error("Wrong bfactor!")
	}
	tree.Insert(5)
	if tree.root.bfactor() != 0 {
		t.Errorf("Failed to LL! root: %v", tree.root)
	}
}

// RL case
func TestRightLeftRotate(t *testing.T) {
	tree := NewTree(8)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(10)
	tree.Insert(14)
	if tree.root.height != 2 {
		t.Error("Failed to insert node!")
	}
	if tree.root.bfactor() != -1 {
		t.Error("Wrong bfactor!")
	}
	tree.Insert(9)
	if tree.root.bfactor() != 0 {
		t.Errorf("Failed to LL! root: %v", tree.root)
	}
}

func TestSearchNode(t *testing.T) {
	tree := NewTree(8)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(10)
	tree.Insert(14)
	tree.Insert(9)
	tree.Insert(5)
	if !tree.SearchNode(14) {
		t.Errorf("Failed to find node %d", 14)
	}
	if tree.SearchNode(0) {
		t.Error("There are no node 0!\n")
	}
}

func TestDeleteNode(t *testing.T) {
	tree := NewTree(8)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(5)
	tree.Insert(11)
	tree.Delete(0)
	if tree.SearchNode(0) {
		t.Errorf("Failed to delete node 0\n")
	}
	tree.Delete(10)
	if tree.SearchNode(10) {
		t.Errorf("Failed to delete node %d", 10)
	}
	if tree.root.value != 9 {
		t.Errorf("Failed to delete node %d", 10)
	}
	tree.Delete(12)
	if tree.SearchNode(12) {
		t.Errorf("Failed to delete node %d", 12)
	}
	if !tree.SearchNode(4) {
		t.Errorf("Failed to find node %d", 4)
	}
}

func TestInorder(t *testing.T) {
	tree := NewTree(8)
	tree.Insert(4)
	tree.Insert(12)
	tree.Insert(10)
	tree.Insert(9)
	tree.Insert(5)
	tree.Insert(11)

	order := tree.Inorder()
	want := []int{4, 5, 8, 9, 10, 11, 12}
	for i := range want {
		if order[i] != want[i] {
			t.Errorf("Wrong order, got: %d, want: %d", order[i], want[i])
		}
	}
}

func TestAVLTreeBigData(t *testing.T) {
	tree := NewTree(0)
	for i := 0; i < 10000; i++ {
		tree.Insert(rand.Intn(1000000))
	}

	for i := 0; i < 1000000; i++ {
		tree.SearchNode(rand.Intn(1000000))
	}

	for i := 0; i < 500; i++ {
		tree.Delete(rand.Intn(1000000))
	}

	for i := 0; i < 1000000; i++ {
		tree.SearchNode(rand.Intn(1000000))
	}
}
