package dag

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDigraph_CheckCircle_NoCircle(t *testing.T) {
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
		return arr
	}
	for _, v := range vertices() {
		err = dg.AddVertex(v)
		assert.NoError(t, err, "failed to add %+v", v)
	}
	assert.False(t, dg.CheckCircle(), "failed to check circle")
}

func TestDigraph_CheckCircle_HasCircle(t *testing.T) {
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
		v, _ = NewVertex(ID("F"), []ID{ID("A")})
		arr = append(arr, v)
		return arr
	}
	for _, v := range vertices() {
		err = dg.AddVertex(v)
		assert.NoError(t, err, "failed to add %+v", v)
	}
	assert.True(t, dg.CheckCircle(), "failed to check circle")
	fmt.Printf("circle: %s\n", dg.CirclePath())
}
