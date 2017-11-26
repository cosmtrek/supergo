package dag

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
)

func TestDigraph_AddVertex(t *testing.T) {
	dg, err := New(6)
	assert.NoError(t, err)

	vertices := func() []Vertex {
		var arr []Vertex
		v, _ := NewVertex(ID("A"), []ID{ID("C"), ID("D")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("B"), []ID{ID("D")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("C"), []ID{ID("E")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("D"), []ID{ID("E")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("E"), []ID{ID("F")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("H"), nil)
		arr = append(arr, v)
		return arr
	}
	for _, v := range vertices() {
		err = dg.AddVertex(v)
		assert.NoError(t, err, "failed to add %+v", v)
	}
}

func TestDigraph_TopologicalOrder(t *testing.T) {
	dg, err := New(6)
	assert.NoError(t, err)

	vertices := func() []Vertex {
		var arr []Vertex
		v, _ := NewVertex(ID("A"), []ID{ID("C"), ID("D")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("B"), []ID{ID("D")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("C"), []ID{ID("E")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("D"), []ID{ID("E")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("E"), []ID{ID("F")})
		arr = append(arr, v)
		v, _ = NewVertex(ID("H"), nil)
		arr = append(arr, v)
		return arr
	}
	for _, v := range vertices() {
		err = dg.AddVertex(v)
		assert.NoError(t, err, "failed to add %+v", v)
	}
	fmt.Printf("%+v\n", dg.TopologicalOrder())
}
