package guid

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGenerator(t *testing.T) {
	node := int64(1)
	g, err := NewGenerator(node)
	assert.Nil(t, err)
	assert.Equal(t, g.node, node)
}

func TestGenerate(t *testing.T) {
	node := int64(1)
	g, err := NewGenerator(node)
	assert.Nil(t, err)
	assert.Equal(t, g.node, node)
	var wg sync.WaitGroup
	for i := 0; i < 50000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = g.Generate()
		}()
	}
	wg.Wait()
}
