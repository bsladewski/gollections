package gollections

import (
	"reflect"
)

// listNode represents a single element in a doubly linked list.
type listNode struct {
	value    interface{}
	previous *listNode
	next     *listNode
}

// addBefore inserts an element to the left of this node.
func (n *listNode) addBefore(value interface{}) *listNode {
	e := &listNode{value: value}
	e.next = n
	if n.previous != nil {
		e.previous = n.previous
		n.previous.next = e
	}
	n.previous = e
	return e
}

// addAfter inserts and element to the right of this node.
func (n *listNode) addAfter(value interface{}) *listNode {
	e := &listNode{value: value}
	e.previous = n
	if n.next != nil {
		e.next = n.next
		n.next.previous = e
	}
	n.next = e
	return e
}

// remove removes this element from its neighbors.
// The neighboring elements are linked if they exist.
func (n *listNode) remove() {
	if n.previous != nil {
		n.previous.next = n.next
	}
	if n.next != nil {
		n.next.previous = n.previous
	}
	n.previous = nil
	n.next = nil
}

// linkedList is an implementation of a doubly linked list.
type linkedList struct {
	head   *listNode
	tail   *listNode
	length int
}

// nodeAt retrieves the element at the specified index.
func (l *linkedList) nodeAt(index int) (*listNode, error) {
	if index < 0 || index >= l.length {
		return nil, ErrIndexOutOfBounds
	}
	current := l.head
	for index > 0 {
		current = current.next
		index--
	}
	return current, nil
}

// Add appends new elements to the end of the collection.
func (l *linkedList) Add(values ...interface{}) {
	for _, value := range values {
		if l.length == 0 {
			e := &listNode{value: value}
			l.head = e
			l.tail = e
		} else {
			l.tail.addAfter(value)
			l.tail = l.tail.next
		}
		l.length++
	}
}

// Clear removes all elements from the collection.
func (l *linkedList) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

// Contains checks if the collection contains all specified values.
func (l *linkedList) Contains(values ...interface{}) bool {
	seen := map[interface{}]bool{}
	current := l.head
	for current != nil {
		seen[current.value] = true
		current = current.next
	}
	for _, value := range values {
		if _, ok := seen[value]; !ok {
			return false
		}
	}
	return true
}

// IsEmpty checks if the collection contains no elements.
func (l *linkedList) IsEmpty() bool {
	return l.length == 0
}

// Remove removes all specified values from the collection.
func (l *linkedList) Remove(values ...interface{}) {
	seen := map[interface{}]bool{}
	for _, value := range values {
		seen[value] = true
	}
	current := l.head
	for current != nil {
		e := current
		current = current.next
		if _, ok := seen[e.value]; ok {
			if e == l.head {
				l.head = l.head.next
			}
			if e == l.tail {
				l.tail = l.tail.previous
			}
			e.remove()
			l.length--
		}
	}
}

// Size gets the number of elements in the collection.
func (l *linkedList) Size() int {
	return l.length
}

// ToArray gets an array representation of the collection.
func (l *linkedList) ToArray() []interface{} {
	array := make([]interface{}, l.length)
	current := l.head
	index := 0
	for current != nil {
		array[index] = current.value
		current = current.next
		index++
	}
	return array
}

// IndexOf gets the first occurance of the specified value or -1 if not found.
func (l *linkedList) IndexOf(value interface{}) int {
	current := l.head
	index := 0
	for current != nil {
		if reflect.DeepEqual(value, current.value) {
			return index
		}
		index++
		current = current.next
	}
	return -1
}

// Insert adds elements at the specified index. Can return index not found error.
func (l *linkedList) Insert(index int, values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}
	current, err := l.nodeAt(index)
	if err != nil {
		return err
	}
	current = current.addBefore(values[0])
	if current.next == l.head {
		l.head = current
	}
	l.length++
	for i := 1; i < len(values); i++ {
		current = current.addAfter(values[i])
		l.length++
	}
	return nil
}

// Get retrieves the value of the element at the specified index.
func (l *linkedList) Get(index int) (interface{}, error) {
	e, err := l.nodeAt(index)
	if err != nil {
		return nil, err
	}
	return e.value, nil
}

// RemoveAt removes the element at the specified index.
func (l *linkedList) RemoveAt(index int) error {
	current, err := l.nodeAt(index)
	if err != nil {
		return err
	}
	if current == l.head {
		l.head = l.head.next
	}
	if current == l.tail {
		l.tail = l.tail.previous
	}
	current.remove()
	l.length--
	return nil
}

// Set overwrites the value of the element at the specified index.
func (l *linkedList) Set(index int, value interface{}) error {
	current, err := l.nodeAt(index)
	if err != nil {
		return err
	}
	current.value = value
	return nil
}

// PeekFirst gets the value of the first element in the collection.
func (l *linkedList) PeekFirst() (interface{}, error) {
	if l.head == nil {
		return nil, ErrNoSuchElement
	}
	return l.head.value, nil
}

// PopFirst gets the value of the first element in the collection. The element is removed.
func (l *linkedList) PopFirst() (interface{}, error) {
	if l.head == nil {
		return nil, ErrNoSuchElement
	}
	temp := l.head
	l.head = l.head.next
	temp.remove()
	l.length--
	return temp.value, nil
}

// PeekLast gets the value of the last element in the collection.
func (l *linkedList) PeekLast() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrNoSuchElement
	}
	return l.tail.value, nil
}

// PopLast gets the value of the last element in the collection. The element is removed.
func (l *linkedList) PopLast() (interface{}, error) {
	if l.tail == nil {
		return nil, ErrNoSuchElement
	}
	temp := l.tail
	l.tail = l.tail.previous
	temp.remove()
	l.length--
	return temp.value, nil
}

// AddFirst adds new elements to the beginning of the collection.
func (l *linkedList) AddFirst(values ...interface{}) {
	if len(values) == 0 {
		return
	}
	for value := range values {
		if l.length == 0 {
			l.Add(value)
			continue
		}
		l.Insert(0, value)
	}
}

// NewLinkedCollection initializes a collection backed by a linked list.
func NewLinkedCollection() Collection {
	return &linkedList{}
}

// NewLinkedList initializes a list backed by a linked list.
func NewLinkedList() List {
	return &linkedList{}
}

// NewLinkedQueue initializes a queue backed by a linked list.
func NewLinkedQueue() Queue {
	return &linkedList{}
}

// NewLinkedStack initializes a stack backed by a linked list.
func NewLinkedStack() Stack {
	return &linkedList{}
}

// NewLinkedDeque initializes a deque backed by a linked list.
func NewLinkedDeque() Deque {
	return &linkedList{}
}
