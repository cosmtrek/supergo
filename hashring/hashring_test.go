package hashring

import (
	"testing"
)

func TestNewHashring(t *testing.T) {
	hr := NewHashring(1, nil)
	if hr.replicas != 1 {
		t.Error("Failed to initialize hash ring.")
	}
}

func TestNodeKey(t *testing.T) {
	n := NewNode(1)
	if n.key(0) != "1:0" {
		t.Error("Wrong key for node!")
	}
}

func TestHashringAddNode(t *testing.T) {
	hr := NewHashring(1, nil)

	nodes := make([]Node, 10)
	for i, v := range nodes {
		v = NewNode(i)
		hr.AddNode(v)
	}

	// Calculated based on https://github.com/redis/redis-rb/blob/master/lib/redis/hash_ring.rb
	//
	// ruby code:
	//   nodes = []
	//   (0...10).each{ |i| nodes << Node.new(i) }
	//   puts Redis::HashRing.new(nodes, 1).sorted_keys.join(",")
	//
	keys := []int{
		3595142127,
		3616270808,
		3629830743,
		3650664544,
		3659295758,
		3688760889,
		3705122533,
		3709191378,
		3734048956,
		3746790027,
	}

	for i, v := range keys {
		if hr.keys[i] != key(v) {
			t.Errorf("Expected key: %d, actual: %d", v, hr.keys[i])
		}
	}
}

func TestHashringRemoveNode(t *testing.T) {
	hr := NewHashring(1, nil)
	nodes := make([]Node, 10)

	for i, v := range nodes {
		v = NewNode(i)
		hr.AddNode(v)
	}

	hr.RemoveNode(nodes[0])

	if len(hr.keys) != 9 {
		t.Errorf("Failed to delete node, keys count: %d", len(hr.keys))
	}
}

func TestHashringGetNode(t *testing.T) {
	hr := NewHashring(0, nil)

	nodes := make([]Node, 10)
	for i, v := range nodes {
		v = NewNode(i)
		hr.AddNode(v)
	}

	keys := []string{
		"port",
		"redis-rb",
		"in",
		"golang",
		"tada",
	}

	retNodes := make([]Node, 0)
	for _, v := range keys {
		if v, ok := hr.GetNode(v); ok {
			retNodes = append(retNodes, v)
		}
	}

	// Calculated based on https://github.com/redis/redis-rb/blob/master/lib/redis/hash_ring.rb
	//
	// ruby code:
	//   nodes = []
	//   (0...10).each{ |i| nodes << Node.new(i) }
	//   hr = Redis::HashRing.new(nodes)
	//   keys = %w(port redis-rb in golang tada)
	//   ret_nodes = []
	//   keys.each{ |k| ret_nodes << hr.get_node(k) }
	//   puts ret_nodes.map(&:id).join(",")
	//
	nodeIDs := []int{7, 4, 4, 5, 6}
	for i, v := range nodeIDs {
		if retNodes[i].ID != v {
			t.Errorf("Expected node id: %d, actual: %d", v, retNodes[i].ID)
		}
	}
}
