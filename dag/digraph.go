package dag

import (
	"container/list"
	"errors"
	"fmt"
)

// Digraph represents a directed graph
type Digraph struct {
	v           int      // number of vertices
	e           int      // number of edges
	vertices    []Vertex // vertices
	verticesMap map[ID]Vertex
	adj         map[ID]*list.List // adj[v] = adjacency list for vertex v
	indegree    map[ID]int        // indegree[v] = indegree of vertex v
	dc          *directedCycle
}

// ID represents node id
type ID string

// New initializes a directed graph
func New(v int) (*Digraph, error) {
	if v < 0 {
		return nil, errors.New("Number of vertices in a Digraph must be nonnegative")
	}
	d := new(Digraph)
	d.verticesMap = make(map[ID]Vertex)
	d.adj = make(map[ID]*list.List, v)
	d.indegree = make(map[ID]int)
	return d, nil
}

// AddVertex adds new vertex
func (d *Digraph) AddVertex(v Vertex) error {
	if !v.isValid() {
		return errors.New("Invalid vertice")
	}
	if _, ok := d.adj[v.ID]; !ok {
		d.v++
		d.vertices = append(d.vertices, v)
		d.verticesMap[v.ID] = v
		d.adj[v.ID] = list.New()
	}
	if v.hasNext() {
		for i := 0; i < len(v.Next); i++ {
			nID := v.Next[i]
			d.e++
			d.adj[v.ID].PushBack(nID)
			d.indegree[nID]++
			nv, err := NewVertex(nID, nil)
			if err != nil {
				return err
			}
			if err = d.AddVertex(nv); err != nil {
				return err
			}
		}
	}
	return nil
}

// CheckCircle return true if circle is found in the graph
func (d *Digraph) CheckCircle() bool {
	d.dc = newDirectedCycle(d)
	return d.dc.checkCircle()
}

// CirclePath returns the circle or empty string if it is not found
func (d *Digraph) CirclePath() string {
	return d.dc.circlePath()
}

// V returns number of vertices
func (d *Digraph) V() int {
	return d.v
}

// E returns number of edges
func (d *Digraph) E() int {
	return d.e
}

// Vertices return all vertices
func (d *Digraph) Vertices() []Vertex {
	return d.vertices
}

// Vertex fetches the vertex by id
func (d *Digraph) Vertex(id ID) Vertex {
	if v, ok := d.verticesMap[id]; ok {
		return v
	}
	return Vertex{ID: ""}
}

// TopologicalOrder returns vertices in topological order
func (d *Digraph) TopologicalOrder() []Vertex {
	if d.CheckCircle() {
		return nil
	}
	return newDepthFirstOrder(d).order()
}

// String inspects graph
func (d *Digraph) String() string {
	ve := fmt.Sprintf("%d vertices, %d edges", d.v, d.e)
	var s string
	for id, edges := range d.adj {
		s += fmt.Sprintf("%s: ", id)
		for e := edges.Front(); e != nil; e = e.Next() {
			s += fmt.Sprintf("-> %s", e.Value)
		}
		s += "\n"
	}
	return fmt.Sprintf("%s\n%s", ve, s)
}

// Vertex holds it's id and id of next vertex if it has
type Vertex struct {
	ID   ID
	Next []ID
}

// NewVertex creates a vertex
func NewVertex(id ID, next []ID) (Vertex, error) {
	return Vertex{
		ID:   id,
		Next: next,
	}, nil
}

func (v Vertex) hasNext() bool {
	return len(v.Next) > 0
}

func (v Vertex) isValid() bool {
	if v.ID == "" {
		return false
	}
	return true
}
