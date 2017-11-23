package avl

// Tree holds the root node.
type Tree struct {
	root *Node
}

// Node contains necessary info.
type Node struct {
	height int
	value  int
	left   *Node
	right  *Node
}

// NewTree initializes an  tree.
func NewTree(value int) *Tree {
	return &Tree{
		root: &Node{
			height: 0,
			value:  value,
			left:   nil,
			right:  nil,
		},
	}
}

// NewNode creates a new  tree node.
func NewNode(value int, left *Node, right *Node) *Node {
	return &Node{
		height: 0,
		value:  value,
		left:   left,
		right:  right,
	}
}

func height(n *Node) int {
	if n == nil {
		return 0
	}
	return n.height
}

func (n *Node) bfactor() int {
	return height(n.left) - height(n.right)
}

func (n *Node) fixHeight() {
	if height(n.left) > height(n.right) {
		n.height = height(n.left) + 1
	} else {
		n.height = height(n.right) + 1
	}
}

// LL case
func (n *Node) rightRotate() *Node {
	m := n.left
	n.left = m.right
	m.right = n
	n.fixHeight()
	m.fixHeight()
	return m
}

// RR case
func (n *Node) leftRotate() *Node {
	m := n.right
	n.right = m.left
	m.left = n
	n.fixHeight()
	m.fixHeight()
	return m
}

// LR case
func (n *Node) leftRightRotate() *Node {
	n.left = n.left.leftRotate()
	return n.rightRotate()
}

// RL case
func (n *Node) rightLeftRotate() *Node {
	n.right = n.right.rightRotate()
	return n.leftRotate()
}

// Insert adds nodes to the tree.
func (t *Tree) Insert(value int) {
	if t == nil {
		return
	}
	n := NewNode(value, nil, nil)
	t.root = insertHelper(t.root, n)
}

func insertHelper(t *Node, n *Node) *Node {
	if t == nil {
		return n
	}

	if n.value < t.value { // left
		t.left = insertHelper(t.left, n)
		if t.bfactor() == 2 {
			if n.value < t.left.value {
				t = t.rightRotate()
			} else {
				t = t.leftRightRotate()
			}
		}
	} else if n.value > t.value { // right
		t.right = insertHelper(t.right, n)
		if t.bfactor() == -2 {
			if n.value > t.right.value {
				t = t.leftRotate()
			} else {
				t = t.rightLeftRotate()
			}
		}
	} else { // equal
		return nil
	}

	t.fixHeight()

	return t
}

// Delete removes a node from the tree.
func (t *Tree) Delete(value int) {
	if t == nil {
		return
	}

	if t.SearchNode(value) {
		node := NewNode(value, nil, nil)
		t.root = deleteNode(t.root, node)
	}
}

// Attention: always return the root node.
func deleteNode(t *Node, n *Node) *Node {
	if t == nil || n == nil {
		return nil
	}

	if n.value < t.value { // left
		t.left = deleteNode(t.left, n)
		if t.bfactor() == -2 {
			r := t.right
			if height(r.left) > height(r.right) {
				t = t.rightLeftRotate()
			} else {
				t = t.leftRotate()
			}
		}
	} else if n.value > t.value { // right
		t.right = deleteNode(t.right, n)
		if t.bfactor() == 2 {
			l := t.left
			if height(l.right) > height(l.left) {
				t = t.leftRightRotate()
			} else {
				t = t.rightRotate()
			}
		}
	} else { // find the node
		if t.left != nil && t.right != nil { // the node has two children
			if height(t.left) > height(t.right) {
				max := maximumNode(t.left)
				t.value = max.value
				t.left = deleteNode(t.left, max)
			} else {
				min := minimumNode(t.right)
				t.value = min.value
				t.right = deleteNode(t.right, min)
			}
		} else { // the node has one child or no children
			if t.left != nil {
				t = t.left
			} else {
				t = t.right
			}
		}
	}

	return t
}

func maximumNode(n *Node) *Node {
	if n == nil {
		return nil
	}

	m := n
	for ; m.right != nil; m = m.right {
	}

	return m
}

func minimumNode(n *Node) *Node {
	if n == nil {
		return nil
	}

	m := n
	for ; m.left != nil; m = m.left {
	}

	return m
}

// SearchNode finds the node.
func (t *Tree) SearchNode(value int) bool {
	return searchNodeHelper(t.root, value)
}

func searchNodeHelper(t *Node, value int) bool {
	if t == nil {
		return false
	}
	if value < t.value {
		return searchNodeHelper(t.left, value)
	} else if value > t.value {
		return searchNodeHelper(t.right, value)
	} else {
		return true
	}
}

// Inorder traverses the tree in middle order.
func (t *Tree) Inorder() []int {
	var result []int
	inorderHelper(t.root, &result)
	return result
}

func inorderHelper(n *Node, result *[]int) {
	if n == nil {
		return
	}

	inorderHelper(n.left, result)
	*result = append(*result, n.value)
	inorderHelper(n.right, result)
}
