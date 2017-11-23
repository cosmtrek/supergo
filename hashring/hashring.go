// Based on https://github.com/redis/redis-rb/blob/master/lib/redis/hash_ring.rb

package hashring

import (
	"fmt"
	"hash/crc32"
	"sort"
)

type key uint32

// Hash ...
type Hash func(key []byte) uint32

// Node ...
type Node struct {
	ID int
}

// Hashring ...
type Hashring struct {
	replicas int
	nodes    []Node
	hash     Hash
	ring     map[key]Node
	keys     []key // sorted
}

const (
	pointPerServer = 160
)

// NewHashring ...
func NewHashring(replicas int, hash Hash) *Hashring {
	h := &Hashring{
		replicas: replicas,
		hash:     hash,
		ring:     make(map[key]Node),
		nodes:    make([]Node, 0),
		keys:     make([]key, 0),
	}

	if h.hash == nil {
		h.hash = crc32.ChecksumIEEE
	}

	if h.replicas <= 0 {
		h.replicas = pointPerServer
	}

	return h
}

// NewNode ...
func NewNode(id int) Node {
	return Node{
		ID: id,
	}
}

func (n Node) key(i int) string {
	return fmt.Sprintf("%d:%d", n.ID, i)
}

// Keys ...
type Keys []key

// Len ...
func (ks Keys) Len() int { return len(ks) }

// Less ...
func (ks Keys) Less(i, j int) bool { return ks[i] < ks[j] }

// Swap ...
func (ks Keys) Swap(i, j int) { ks[i], ks[j] = ks[j], ks[i] }

// AddNode ...
func (h *Hashring) AddNode(n Node) {
	h.nodes = append(h.nodes, n)

	for i := 0; i < h.replicas; i++ {
		k := key(h.hash([]byte(n.key(i))))
		h.ring[k] = n
		h.keys = append(h.keys, k)
	}

	sort.Stable(Keys(h.keys))
}

// RemoveNode ...
func (h *Hashring) RemoveNode(n Node) {
	for i, v := range h.nodes {
		if v.ID == n.ID {
			h.nodes = append(h.nodes[:i], h.nodes[i+1:]...)
		}
	}

	for i := 0; i < h.replicas; i++ {
		k := key(h.hash([]byte(n.key(i))))
		delete(h.ring, k)

		for i, v := range h.keys {
			if v == k {
				h.keys = append(h.keys[:i], h.keys[i+1:]...)
			}
		}
	}

	sort.Stable(Keys(h.keys))
}

// IsEmpty ...
func (h *Hashring) IsEmpty() bool {
	return len(h.keys) == 0
}

// GetNode ...
func (h *Hashring) GetNode(data string) (Node, bool) {
	if h.IsEmpty() {
		return NewNode(-1), false
	}

	k := key(h.hash([]byte(data)))

	// binary search for appropriate replica
	idx := sort.Search(len(h.keys), func(i int) bool { return h.keys[i] >= k })
	if idx <= 0 {
		idx = len(h.keys) - 1
	} else {
		idx--
	}

	return h.ring[h.keys[idx]], true
}
