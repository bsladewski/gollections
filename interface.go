package gollections

import "errors"

// A Collection is a grouping of elements.
type Collection interface {
	// Add appends new elements to the end of the collection.
	Add(values ...interface{})
	// Clear removes all elements from the collection.
	Clear()
	// Contains checks if the collection contains all specified values.
	Contains(values ...interface{}) bool
	// IsEmpty checks if the collection contains no elements.
	IsEmpty() bool
	// Remove removes all specified values from the collection.
	Remove(values ...interface{})
	// Size gets the number of elements in the collection.
	Size() int
	// ToArray gets an array representation of the collection.
	ToArray() []interface{}
}

// ErrIndexOutOfBounds the supplied index was invalid for this list.
var ErrIndexOutOfBounds = errors.New("index out of bounds")

// A List is an ordered collection that can be accessed by index.
type List interface {
	Collection
	// IndexOf gets the first occurance of the specified value or -1 if not found.
	IndexOf(value interface{}) int
	// Insert adds elements at the specified index. Can return index not found error.
	Insert(index int, values ...interface{}) error
	// Get retrieves the value of the element at the specified index.
	Get(index int) (interface{}, error)
	// RemoveAt removes the element at the specified index.
	RemoveAt(index int) error
	// Set overwrites the value of the element at the specified index.
	Set(index int, value interface{}) error
}

// ErrNoSuchElement the polled element does not exist.
var ErrNoSuchElement = errors.New("no such element")

// A Queue provides FIFO access to a collection.
type Queue interface {
	Collection
	// PeekFirst gets the value of the first element in the collection.
	PeekFirst() (interface{}, error)
	// PopFirst gets the value of the first element in the collection. The element is removed.
	PopFirst() (interface{}, error)
}

// A Deque is a double ended queue.
type Deque interface {
	Queue
	// AddFirst adds new elements to the beginning of the collection.
	AddFirst(values ...interface{})
	// PeekLast gets the value of the last element in the collection.
	PeekLast() (interface{}, error)
	// PopLast gets the value of the last element in the collection. The element is removed.
	PopLast() (interface{}, error)
}

// A Stack provides FILO/LIFO access to a collection.
type Stack interface {
	Collection
	// PeekLast gets the value of the last element in the collection.
	PeekLast() (interface{}, error)
	// PopLast gets the value of the last element in the collection. The element is removed.
	PopLast() (interface{}, error)
}
