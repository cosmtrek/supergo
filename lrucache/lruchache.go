package lrucache

import (
	"container/list"
	"sync"
)

// Entry ...
type Entry struct {
	key string
	val interface{}
}

// LRUCache ...
type LRUCache struct {
	maxSize   int
	entryList *list.List
	entryMap  map[string]*list.Element
	sync.RWMutex
}

func newLRUCache(maxSize int) *LRUCache {
	if maxSize <= 0 {
		return nil
	}
	return &LRUCache{
		maxSize:   maxSize,
		entryList: list.New(),
		entryMap:  make(map[string]*list.Element, 0),
	}
}

// Get ...
func (c *LRUCache) Get(key string) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()
	entry, ok := c.entryMap[key]
	if !ok {
		return nil, false
	}

	c.entryList.MoveToFront(entry)
	e := entry.Value.(*Entry)
	return e.val, true
}

// Set ...
func (c *LRUCache) Set(key string, val interface{}) *LRUCache {
	c.Lock()
	defer c.Unlock()
	entry, ok := c.entryMap[key]
	if ok {
		c.entryList.PushFront(entry)
		e := entry.Value.(*Entry)
		e.val = val
	} else {
		entry = c.entryList.PushFront(&Entry{key, val})
		c.entryMap[key] = entry

		if c.entryList.Len() > c.maxSize {
			removedEntry := c.entryList.Back()
			k := removedEntry.Value.(*Entry)
			c.entryList.Remove(removedEntry)
			delete(c.entryMap, k.key)
		}
	}
	return c
}

// MaxSize ...
func (c *LRUCache) MaxSize() int {
	c.RLock()
	defer c.RUnlock()

	if c == nil {
		return 0
	}
	return c.maxSize
}

// Size ...
func (c *LRUCache) Size() int {
	c.RLock()
	defer c.RUnlock()

	if c == nil {
		return 0
	}
	return c.entryList.Len()
}
