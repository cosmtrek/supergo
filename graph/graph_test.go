package graph

import (
	"testing"
)

func compareArray(t *testing.T, got, want []int) {
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("Got: %d, expected: %d", got[i], want[i])
		}
	}
}

func TestNewUGraph(t *testing.T) {
	v, e := exampleGraph()
	g := NewUGraph(v, e)
	if g == nil {
		t.Error("Failed to new graph!")
	}
	if g.matrix[0][2] != 1 || g.matrix[2][0] != 1 {
		t.Error("Failed to create an edge A<->C")
	}
	if g.matrix[5][6] != 1 || g.matrix[6][5] != 1 {
		t.Error("Failed to create an edge F<->G")
	}
}

func exampleGraph() ([]byte, [][]byte) {
	vexs := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	edges := [][]byte{
		{'A', 'C'},
		{'A', 'D'},
		{'A', 'F'},
		{'B', 'C'},
		{'C', 'D'},
		{'E', 'G'},
		{'F', 'G'},
	}
	return vexs, edges
}

func Graph() *UGraph {
	v, e := exampleGraph()
	return NewUGraph(v, e)
}

func TestDFS(t *testing.T) {
	g := Graph()
	path := g.DFS()
	want := []int{0, 2, 1, 3, 5, 6, 4}
	compareArray(t, path, want)
}

func TestBFS(t *testing.T) {
	g := Graph()
	path := g.BFS()
	want := []int{0, 2, 3, 5, 1, 6, 4}
	compareArray(t, path, want)
}

func exampleUWGraph() ([]byte, [][]int) {
	vexs := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G'}
	matrix := [][]int{
		{0, 12, MaxInt, MaxInt, MaxInt, 16, 14},
		{12, 0, 10, MaxInt, MaxInt, 7, MaxInt},
		{MaxInt, 10, 0, 3, 5, 6, MaxInt},
		{MaxInt, MaxInt, 3, 0, 4, MaxInt, MaxInt},
		{MaxInt, MaxInt, 5, 4, 0, 2, 8},
		{16, 7, 6, MaxInt, 2, 0, 9},
		{14, MaxInt, MaxInt, MaxInt, 8, 9, 0},
	}
	return vexs, matrix
}

func WGraph() *UWGraph {
	v, e := exampleUWGraph()
	return NewUWGraph(v, e)
}

func TestNewUWGraph(t *testing.T) {
	g := WGraph()
	if g == nil {
		t.Error("Failed to new a weighted graph")
	}
	if g.matrix[0][1] != 12 && g.matrix[1][0] != 12 {
		t.Errorf("Got: %d, want: %d", g.matrix[0][1], 12)
	}
	if g.matrix[1][3] != MaxInt {
		t.Errorf("Got: %d, want: %d", g.matrix[1][3], MaxInt)
	}
}

func TestDijkstra(t *testing.T) {
	g := WGraph()
	// A is start point
	want := []int{0, 12, 22, 22, 18, 16, 14}
	got := g.Dijkstra(0)
	compareArray(t, got, want)
}
