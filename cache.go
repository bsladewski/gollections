package gollections

import (
	"sync"
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

// entry holds a pointer to a node in the key list and the cached value.
type entry struct {
	node  *listNode
	value interface{}
}

// A concurrentCache provides fast, thread safe access to key/value pairs.
type concurrentCache struct {
	maxLength int
	values    map[interface{}]entry
	keys      *LinkedList
	mutex     *sync.RWMutex
}

// prune removes elements from the head of the cache value list.
// As elements are added to the tail when added or accessed, the head will always contain the
// oldest entry in the cache.
func (cache *concurrentCache) prune() error {
	if cache.maxLength <= 0 {
		return nil
	}
	for cache.keys.length > cache.maxLength {
		key, err := cache.keys.Get(0)
		if err != nil {
			return err
		}
		cache.keys.RemoveAt(0)
		delete(cache.values, key)
	}
	return nil
}

// touch moves an existing node to the tail of the key list.
func (cache *concurrentCache) touch(node *listNode) {
	if node != cache.keys.tail {
		node.remove()
		cache.keys.tail.next = node
		node.previous = cache.keys.tail
		cache.keys.tail = node
	}
}

func (cache *concurrentCache) Clear() {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.keys.Clear()
	cache.values = map[interface{}]entry{}
}

func (cache *concurrentCache) Get(key interface{}) (interface{}, error) {
	cache.mutex.RLock()
	defer cache.mutex.RUnlock()
	if entry, ok := cache.values[key]; ok {
		cache.touch(entry.node)
		return entry.value, nil
	}
	return nil, ErrNoSuchElement
}

func (cache *concurrentCache) Put(key interface{}, value interface{}) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if entry, ok := cache.values[key]; ok {
		cache.touch(entry.node)
	} else {
		cache.keys.Add(key)
	}
	cache.values[key] = entry{node: cache.keys.tail, value: value}
	cache.prune()
}

func (cache *concurrentCache) SetMaxSize(maxSize int) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.maxLength = maxSize
	cache.prune()
}

func (cache *concurrentCache) Size() int {
	cache.mutex.RLock()
	cache.mutex.RUnlock()
	return cache.keys.length
}

func (cache *concurrentCache) Remove(key interface{}) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	if entry, ok := cache.values[key]; ok {
		if entry.node == cache.keys.head {
			cache.keys.head = entry.node.next
		}
		if entry.node == cache.keys.tail {
			cache.keys.tail = entry.node.previous
		}
		entry.node.remove()
		cache.keys.length--
		delete(cache.values, key)
	}
}
