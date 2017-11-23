package rbtree

// Tree ...
type Tree struct {
	root *Node
}

type color int

// Color
const (
	RED color = iota
	BLACK
)

// Node ...
type Node struct {
	color  color
	key    int
	left   *Node
	right  *Node
	parent *Node
}

// NewTree ...
func NewTree() *Tree {
	return &Tree{
		root: nil,
	}
}

// NewNode ...
func NewNode(key int) *Node {
	return &Node{
		key:    key,
		color:  BLACK,
		left:   nil,
		right:  nil,
		parent: nil,
	}
}

func (t *Tree) leftRotate(n *Node) {
	m := n.right
	n.right = m.left
	if m.left != nil {
		m.left.parent = n
	}
	m.parent = n.parent

	if n.parent == nil {
		t.root = m
	} else {
		if n.parent.left == n {
			n.parent.left = m
		} else {
			n.parent.right = m
		}
	}

	m.left = n
	n.parent = m
}

func (t *Tree) rightRotate(n *Node) {
	m := n.left
	n.left = m.right
	if m.right != nil {
		m.right.parent = n
	}
	m.parent = n.parent

	if n.parent == nil {
		t.root = m
	} else {
		if n.parent.right == n {
			n.parent.right = m
		} else {
			n.parent.left = m
		}
	}

	m.right = n
	n.parent = m
}

// Insert adds new node to tree.
func (t *Tree) Insert(key int) {
	if _, ok := t.Search(key); ok {
		return
	}

	n := NewNode(key)
	var y *Node
	x := t.root

	for x != nil {
		y = x
		if n.key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}
	n.parent = y

	if y != nil {
		if n.key < y.key {
			y.left = n
		} else {
			y.right = n
		}
	} else {
		t.root = n
	}

	n.color = RED
	t.insertFixup(n)
}

func (t *Tree) insertFixup(n *Node) {
	var p, u, gp *Node

	for n.parent != nil && n.parent.color == RED {
		p = n.parent
		gp = p.parent

		// parent node is left child
		if p == gp.left {
			// case 1: uncle node is red
			u = gp.right
			if u != nil && u.color == RED {
				u.color = BLACK
				p.color = BLACK
				gp.color = RED
				n = gp
				continue
			}

			// case 2: uncle node is black and right child
			if p.right == n {
				t.leftRotate(p)
				tmp := p
				p = n
				n = tmp
			}

			// case 3: uncle node is black and left child
			p.color = BLACK
			gp.color = RED
			t.rightRotate(gp)
		} else { // parent node is right child
			// case 1: uncle node is red
			u = gp.left
			if u != nil && u.color == RED {
				u.color = BLACK
				p.color = BLACK
				gp.color = RED
				n = gp
				continue
			}

			// case 2: uncle node is black and left child
			if p.left == n {
				t.rightRotate(p)
				tmp := p
				p = n
				n = tmp
			}

			// case 3: uncle node is black and right child
			p.color = BLACK
			gp.color = RED
			t.leftRotate(gp)
		}
	}

	t.root.color = BLACK
}

// Delete ...
func (t *Tree) Delete(key int) {
	n, ok := t.Search(key)
	if !ok {
		return
	}

	var c, p *Node
	var clr color

	// n has two children
	if n.left != nil && n.right != nil {
		next := n
		next = n.right
		for next.left != nil {
			next = next.left
		}

		if n.parent != nil {
			if n.parent.left == n {
				n.parent.left = next
			} else {
				n.parent.right = next
			}
		} else {
			t.root = next
		}

		c = next.right
		p = next.parent
		clr = next.color

		if p == n {
			p = next
		} else {
			if c != nil {
				c.parent = p
			}
			p.left = c

			next.right = n.right
			n.right.parent = next
		}

		next.parent = n.parent
		next.color = n.color
		next.left = n.left
		n.left.parent = next

		if clr == BLACK {
			t.deleteFixup(c, p)
		}

		return
	}

	if n.left != nil {
		c = n.left
	} else {
		c = n.right
	}

	p = n.parent
	clr = n.color

	if c != nil {
		c.parent = p
	}

	if p != nil {
		if p.left == n {
			p.left = c
		} else {
			p.right = c
		}
	} else {
		t.root = c
	}

	if clr == BLACK {
		t.deleteFixup(c, p)
	}
}

func (t *Tree) deleteFixup(n *Node, p *Node) {
	var other *Node

	for n != nil && n.color == BLACK && n != t.root {
		if p.left == n { // left child
			other = p.right
			if other.color == RED {
				// case 1
				other.color = BLACK
				p.color = RED
				t.leftRotate(p)
				other = p.right
			}

			if (other.left != nil || other.left.color == BLACK) &&
				(other.right != nil || other.right.color == BLACK) {
				// case 2
				other.color = RED
				n = p
				p = n.parent
			} else {
				if other.right != nil || other.right.color == BLACK {
					// case 3
					other.left.color = BLACK
					other.color = RED
					t.rightRotate(other)
					other = p.right
				}
				// case 4
				other.color = p.color
				p.color = BLACK
				other.right.color = BLACK
				t.leftRotate(p)
				n = t.root
				break
			}
		} else { // right child
			other = p.left
			if other.color == RED {
				// case 1
				other.color = BLACK
				p.color = RED
				t.rightRotate(p)
				other = p.left
			}

			if (other.right != nil || other.right.color == BLACK) &&
				(other.left != nil || other.left.color == BLACK) {
				// case 2
				other.color = RED
				n = p
				p = n.parent
			} else {
				if other.left != nil || other.left.color == BLACK {
					// case 3
					other.right.color = BLACK
					other.color = RED
					t.leftRotate(other)
					other = p.left
				}
				// case 4
				other.color = p.color
				p.color = BLACK
				other.left.color = BLACK
				t.rightRotate(p)
				n = t.root
				break
			}
		}
	}

	if n != nil {
		n.color = BLACK
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
	*result = append(*result, n.key)
	inorderHelper(n.right, result)
}

// Search ...
func (t *Tree) Search(key int) (*Node, bool) {
	if t == nil {
		return nil, false
	}

	x := t.root
	for x != nil && x.key != key {
		if key < x.key {
			x = x.left
		} else {
			x = x.right
		}
	}

	if x == nil {
		return nil, false
	}
	return x, true
}
