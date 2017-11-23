package graph

// UGraph is a matrix undirected graph.
type UGraph struct {
	vex    int
	edge   int
	vexs   []byte  // vertex
	matrix [][]int // edge
}

// Edge stores an edge info.
type Edge struct {
	start  byte
	end    byte
	weight int
}

// UWGraph is a matrix undirected weighted graph.
type UWGraph struct {
	UGraph
}

// NewUGraph ...
func NewUGraph(vexs []byte, edges [][]byte) *UGraph {
	if vexs == nil || edges == nil {
		return nil
	}

	vl := len(vexs)
	g := UGraph{
		vex:    vl,
		edge:   len(edges),
		vexs:   make([]byte, vl),
		matrix: make([][]int, vl),
	}

	for i := range vexs {
		g.vexs[i] = vexs[i]
	}

	for i := 0; i < g.edge; i++ {
		g.matrix[i] = make([]int, g.vex)
	}

	for i := 0; i < g.edge; i++ {
		from := getPosition(g.vexs, edges[i][0])
		to := getPosition(g.vexs, edges[i][1])
		g.matrix[from][to] = 1
		g.matrix[to][from] = 1
	}
	return &g
}

// Max cons
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// NewUWGraph Undirected weighted graph.
func NewUWGraph(vexs []byte, matrix [][]int) *UWGraph {
	if vexs == nil || matrix == nil {
		return nil
	}

	vl := len(vexs)
	g := UWGraph{
		UGraph: UGraph{
			vex:    vl,
			edge:   0,
			vexs:   make([]byte, vl),
			matrix: make([][]int, vl),
		},
	}

	for i := range vexs {
		g.vexs[i] = vexs[i]
	}

	for i := 0; i < g.vex; i++ {
		g.matrix[i] = make([]int, g.vex)
	}

	for i := 0; i < g.vex; i++ {
		for j := 0; j < g.vex; j++ {
			g.matrix[i][j] = matrix[i][j]
		}
	}

	for i := 0; i < g.vex; i++ {
		for j := 0; j < g.vex; j++ {
			if i != j && g.matrix[i][j] != MaxInt {
				g.edge++
			}
		}
	}
	g.edge /= 2

	return &g
}

func getPosition(vexs []byte, v byte) int {
	for i := range vexs {
		if v == vexs[i] {
			return i
		}
	}
	return -1
}

// DFS ...
func (g *UGraph) DFS() []int {
	visited := make([]bool, g.vex)
	path := make([]int, 0)

	for i := 0; i < g.vex; i++ {
		if !visited[i] {
			dfsHelper(g, i, &visited, &path)
		}
	}
	return path
}

func dfsHelper(g *UGraph, i int, visited *[]bool, path *[]int) {
	var w int
	(*visited)[i] = true
	*path = append(*path, i)
	for w = g.firstVertex(i); w >= 0; w = g.nextVertex(i, w) {
		if !(*visited)[w] {
			dfsHelper(g, w, visited, path)
		}
	}
}

func (g *UGraph) firstVertex(v int) int {
	if v < 0 || v > (g.vex-1) {
		return -1
	}

	for i := 0; i < g.vex; i++ {
		if g.matrix[v][i] == 1 {
			return i
		}
	}

	return -1
}

func (g *UGraph) nextVertex(v int, w int) int {
	if v < 0 || v > (g.vex-1) || w < 0 || w > (g.vex-1) {
		return -1
	}

	for i := w + 1; i < g.vex; i++ {
		if g.matrix[v][i] == 1 {
			return i
		}
	}

	return -1
}

// BFS ...
func (g *UGraph) BFS() []int {
	var head, rear int // queue front and end
	queue := make([]int, g.vex)
	visited := make([]bool, g.vex)
	path := make([]int, 0)

	for i := 0; i < g.vex; i++ {
		if !visited[i] {
			visited[i] = true
			queue[rear] = i // enqueue
			rear++
		}

		for head != rear {
			j := queue[head]
			head++
			path = append(path, j)
			for w := g.firstVertex(j); w >= 0; w = g.nextVertex(j, w) {
				if !visited[w] {
					visited[w] = true
					queue[rear] = w
					rear++
				}
			}
		}
	}

	return path
}

// Dijkstra ...
func (g *UWGraph) Dijkstra(v int) []int {
	flag := make([]int, g.vex)
	prev := make([]int, g.vex)
	dist := make([]int, g.vex)

	for i := 0; i < g.vex; i++ {
		dist[i] = g.matrix[v][i]
	}

	flag[v] = 1
	dist[v] = 0
	k := 0

	for i := 1; i < g.vex; i++ {
		min := MaxInt
		for j := 0; j < g.vex; j++ {
			if flag[j] == 0 && dist[j] < min {
				min = dist[j]
				k = j
			}
		}

		flag[k] = 1

		for j := 0; j < g.vex; j++ {
			var tmp int
			if g.matrix[k][j] == MaxInt {
				tmp = MaxInt
			} else {
				tmp = min + g.matrix[k][j]
			}
			if flag[j] == 0 && tmp < dist[j] {
				dist[j] = tmp
				prev[j] = k
			}
		}
	}

	return dist
}
