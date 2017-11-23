package trie

// NodeLink ...
type NodeLink struct {
	value rune
	link  *Node
}

// Node ...
type Node struct {
	children []NodeLink
}

// Trie ...
type Trie struct {
	root *Node
}

// NewTrie ...
func NewTrie() *Trie {
	return &Trie{
		root: newNode(),
	}
}

func newNode() *Node {
	return &Node{
		children: make([]NodeLink, 0),
	}
}

// Add ...
func (t *Trie) Add(s string) {
	if s == "" {
		return
	}
	t.root.addValue(s)
}

func (n *Node) addValue(s string) {
	if n == nil {
		return
	}

	cur := n
	for _, v := range s {
		nn := n.findLink(v)
		if nn == nil {
			nn = newNode()
			cur.children = append(cur.children, NodeLink{value: v, link: nn})
		}
		cur = nn
	}
}

// Found ...
func (t *Trie) Found(s string) bool {
	cur := t.root

	for _, v := range s {
		nn := cur.findLink(v)
		if nn == nil {
			return false
		}
		cur = nn
	}
	return true
}

func (n *Node) findLink(val rune) *Node {
	for _, v := range n.children {
		if val == v.value {
			return v.link
		}
	}
	return nil
}

// Destroy ...
func (t *Trie) Destroy() {
	t.root = newNode()
}

// Empty ...
func (t *Trie) Empty() bool {
	return len(t.root.children) == 0
}
