// Package cache provides a data structure for caching key/value pairs.
package cache

import (
	"sync"

	"github.com/bsladewski/gollections"
)

// A Cache represents a key/value store.
type Cache interface {
	// Clear removes all entries from the cache.
	Clear()
	// Get retrieves a value from the cache. Returns an error if no such entry exists.
	Get(key interface{}) (interface{}, error)
	// Put adds or updates an entry in the cache.
	Put(key interface{}, value interface{})
	// SetMaxSize updates the maximum number of entries allows in the cache.
	SetMaxSize(maxSize int)
	// Size gets the current number of entries in the cache.
	Size() int
	// Remove deletes a single entry from the cache.
	Remove(key interface{})
}

// A cache provides access to key/value pairs.
type cache struct {
	maxSize int
	values  map[interface{}]interface{}
	keys    gollections.List
}

// prune removes elements from the head of the cache value list.
// As elements are added to the tail when added or accessed, the head will always contain the
// oldest entry in the cache.
func (c *cache) prune() error {
	if c.maxSize <= 0 {
		return nil
	}
	for c.keys.Size() > c.maxSize {
		key, err := c.keys.Get(0)
		if err != nil {
			return err
		}
		c.keys.RemoveAt(0)
		delete(c.values, key)
	}
	return nil
}

// touch moves an existing node to the tail of the key list.
func (c *cache) touch(key interface{}) {
	c.keys.Remove(key)
	c.keys.Add(key)
}

func (c *cache) Clear() {
	c.keys.Clear()
	c.values = map[interface{}]interface{}{}
}

func (c *cache) Get(key interface{}) (interface{}, error) {
	if value, ok := c.values[key]; ok {
		c.touch(key)
		return value, nil
	}
	return nil, gollections.ErrNoSuchElement
}

func (c *cache) Put(key interface{}, value interface{}) {
	if _, ok := c.values[key]; ok {
		c.touch(key)
	} else {
		c.keys.Add(key)
	}
	c.values[key] = value
	c.prune()
}

func (c *cache) SetMaxSize(maxSize int) {
	c.maxSize = maxSize
	c.prune()
}

func (c *cache) Size() int {
	return c.keys.Size()
}

func (c *cache) Remove(key interface{}) {
	c.keys.Remove(key)
	delete(c.values, key)
}

// A concurrentCache synchronizes a standard cache using a read/write mutex.
type concurrentCache struct {
	cache
	mutex *sync.RWMutex
}

func (c *concurrentCache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Clear()
}

func (c *concurrentCache) Get(key interface{}) (interface{}, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.cache.Get(key)
}

func (c *concurrentCache) Put(key interface{}, value interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Put(key, value)
}

func (c *concurrentCache) SetMaxSize(maxSize int) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.SetMaxSize(maxSize)
}

func (c *concurrentCache) Size() int {
	c.mutex.RLock()
	c.mutex.RUnlock()
	return c.cache.Size()
}

func (c *concurrentCache) Remove(key interface{}) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Remove(key)
}

// NewCache initializes a new cache.
func NewCache(maxSize int) Cache {
	return &cache{
		maxSize: maxSize,
		values:  map[interface{}]interface{}{},
		keys:    gollections.NewLinkedList(),
	}
}

// NewConcurrentCache initializes a new thead-safe cache.
func NewConcurrentCache(maxSize int) Cache {
	return &concurrentCache{
		cache: cache{
			maxSize: maxSize,
			values:  map[interface{}]interface{}{},
			keys:    gollections.NewLinkedList(),
		},
		mutex: &sync.RWMutex{},
	}
}
