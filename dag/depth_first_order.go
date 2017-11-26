package dag

import (
	"container/list"
)

type depthFirstOrder struct {
	graph     *Digraph
	marked    map[ID]bool
	post      map[ID]int
	postOrder *list.List
}

func newDepthFirstOrder(d *Digraph) *depthFirstOrder {
	return &depthFirstOrder{
		graph:     d,
		marked:    make(map[ID]bool),
		post:      make(map[ID]int),
		postOrder: list.New(),
	}
}

func (o *depthFirstOrder) order() []Vertex {
	for _, v := range o.graph.Vertices() {
		if _, ok := o.marked[v.ID]; ok {
			continue
		}
		o.dfs(v.ID)
	}
	var order []Vertex
	for e := o.postOrder.Back(); e != nil; e = e.Prev() {
		order = append(order, o.graph.Vertex(e.Value.(ID)))
	}
	return order
}

func (o *depthFirstOrder) dfs(id ID) {
	o.marked[id] = true
	for e := o.graph.adj[id].Front(); e != nil; e = e.Next() {
		ev := e.Value.(ID)
		if _, ok := o.marked[ev]; !ok {
			o.dfs(ev)
		}
	}
	o.postOrder.PushBack(id)
}
