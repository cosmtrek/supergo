package dag

import (
	"container/list"
	"fmt"
)

type directedCycle struct {
	graph     *Digraph
	marked    map[ID]bool // marked[v] = has vertex v been marked?
	onStack   map[ID]bool // onStack[v] = is vertex on the stack?
	edgeTo    map[ID]ID   // edgeTo[v] = previous vertex on path to v
	circle    *list.List  // directed cycle (or null if no such cycle)
	hasCircle bool
}

func newDirectedCycle(d *Digraph) *directedCycle {
	return &directedCycle{
		graph:   d,
		marked:  make(map[ID]bool, d.V()),
		onStack: make(map[ID]bool, d.V()),
		edgeTo:  make(map[ID]ID),
		circle:  list.New(),
	}
}

func (c *directedCycle) checkCircle() bool {
	for _, v := range c.graph.Vertices() {
		if c.hasCircle {
			return true
		}
		if _, ok := c.marked[v.ID]; ok {
			continue
		}
		c.dfs(v.ID)
	}
	return c.hasCircle
}

func (c *directedCycle) dfs(id ID) {
	c.marked[id] = true
	c.onStack[id] = true
	for e := c.graph.adj[id].Front(); e != nil; e = e.Next() {
		if c.circle.Len() > 0 {
			c.hasCircle = true
			return
		}
		ev := e.Value.(ID)
		if _, ok := c.marked[ev]; !ok {
			c.edgeTo[ev] = id
			c.dfs(ev)
		}
		if v, ok := c.onStack[ev]; ok && v {
			for x := id; x != ev && x != ""; x = c.edgeTo[x] {
				c.circle.PushBack(x)
			}
			c.circle.PushBack(ev)
			c.circle.PushBack(id)
		}
	}
	c.onStack[id] = false
}

func (c *directedCycle) circlePath() string {
	if !c.hasCircle {
		return ""
	}
	var path string
	for e := c.circle.Back(); e != nil; e = e.Prev() {
		if e.Prev() == nil {
			path += fmt.Sprintf("%s", e.Value)
		} else {
			path += fmt.Sprintf("%s -> ", e.Value)
		}
	}
	return path
}
