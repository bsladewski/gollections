package gollections

import (
	"errors"
	"reflect"
)

// listNode represents a single element in a doubly linked list.
type listNode struct {
	value    interface{}
	previous *listNode
	next     *listNode
}

// addBefore inserts an element to the left of this node.
func (node *listNode) addBefore(value interface{}) *listNode {
	e := &listNode{value: value}
	e.next = node
	if node.previous != nil {
		e.previous = node.previous
		node.previous.next = e
	}
	node.previous = e
	return e
}

// addAfter inserts and element to the right of this node.
func (node *listNode) addAfter(value interface{}) *listNode {
	e := &listNode{value: value}
	e.previous = node
	if node.next != nil {
		e.next = node.next
		node.next.previous = e
	}
	node.next = e
	return e
}

// remove removes this element from its neighbors.
// The neighboring elements are linked if they exist.
func (node *listNode) remove() {
	if node.previous != nil {
		node.previous.next = node.next
	}
	if node.next != nil {
		node.next.previous = node.previous
	}
	node.previous = nil
	node.next = nil
}

// LinkedList is an implementation of a doubly linked list.
type LinkedList struct {
	head   *listNode
	tail   *listNode
	length int
}

// nodeAt retrieves the element at the specified index.
func (list *LinkedList) nodeAt(index int) (*listNode, error) {
	if index < 0 || index >= list.length {
		return nil, errors.New("index out of bounds")
	}
	current := list.head
	for index > 0 {
		current = current.next
		index--
	}
	return current, nil
}

// Add appends new elements to the end of the collection.
func (list *LinkedList) Add(values ...interface{}) {
	for _, value := range values {
		e := &listNode{value: value}
		if list.length == 0 {
			list.head = e
			list.tail = e
		} else {
			list.tail.addAfter(e)
		}
		list.length++
	}
}

// Clear removes all elements from the collection.
func (list *LinkedList) Clear() {
	list.head = nil
	list.tail = nil
	list.length = 0
}

// Contains checks if the collection contains all specified values.
func (list *LinkedList) Contains(values ...interface{}) bool {
	seen := map[interface{}]bool{}
	current := list.head
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
func (list *LinkedList) IsEmpty() bool {
	return list.length == 0
}

// Remove removes all specified values from the collection.
func (list *LinkedList) Remove(values ...interface{}) {
	seen := map[interface{}]bool{}
	for _, value := range values {
		seen[value] = true
	}
	current := list.head
	for current != nil {
		e := current
		current = current.next
		if _, ok := seen[current.value]; ok {
			e.remove()
			list.length--
		}
	}
}

// Size gets the number of elements in the collection.
func (list *LinkedList) Size() int {
	return list.length
}

// ToArray gets an array representation of the collection.
func (list *LinkedList) ToArray() []interface{} {
	array := []interface{}{}
	current := list.head
	for current != nil {
		array = append(array, current.value)
		current = current.next
	}
	return array
}

// IndexOf gets the first occurance of the specified value or -1 if not found.
func (list *LinkedList) IndexOf(value interface{}) int {
	current := list.head
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
func (list *LinkedList) Insert(index int, values ...interface{}) error {
	if len(values) == 0 {
		return nil
	}
	current, err := list.nodeAt(index)
	if err != nil {
		return err
	}
	current = current.addBefore(values[0])
	for i := 1; i < len(values); i++ {
		current = current.addAfter(values[i])
	}
	return nil
}

// Get retrieves the value of the element at the specified index.
func (list *LinkedList) Get(index int) (interface{}, error) {
	e, err := list.nodeAt(index)
	if err != nil {
		return nil, err
	}
	return e.value, nil
}

// RemoveAt removes the element at the specified index.
func (list *LinkedList) RemoveAt(index int) error {
	current, err := list.nodeAt(index)
	if err != nil {
		return err
	}
	current.remove()
	return nil
}

// Set overwrites the value of the element at the specified index.
func (list *LinkedList) Set(index int, value interface{}) error {
	current, err := list.nodeAt(index)
	if err != nil {
		return err
	}
	current.value = value
	return nil
}

// PeekFirst gets the value of the first element in the collection.
func (list *LinkedList) PeekFirst() (interface{}, error) {
	if list.head == nil {
		return nil, errors.New("no such element")
	}
	return list.head.value, nil
}

// PopFirst gets the value of the first element in the collection. The element is removed.
func (list *LinkedList) PopFirst() (interface{}, error) {
	if list.head == nil {
		return nil, errors.New("no such element")
	}
	temp := list.head
	list.head = list.head.next
	temp.remove()
	return temp.value, nil
}

// PeekLast gets the value of the last element in the collection.
func (list *LinkedList) PeekLast() (interface{}, error) {
	if list.tail == nil {
		return nil, errors.New("no such element")
	}
	return list.tail.value, nil
}

// PopLast gets the value of the last element in the collection. The element is removed.
func (list *LinkedList) PopLast() (interface{}, error) {
	if list.tail == nil {
		return nil, errors.New("no such element")
	}
	temp := list.tail
	list.tail = list.tail.previous
	temp.remove()
	return temp.value, nil
}

// AddFirst adds new elements to the beginning of the collection.
func (list *LinkedList) AddFirst(values ...interface{}) {
	if len(values) == 0 {
		return
	}
	if list.length == 0 {
		list.head = &listNode{value: values[len(values)-1]}
		list.tail = list.head
		list.length = 1
		list.AddFirst(values[:len(values)-1])
	} else {
		list.Insert(0, values)
	}
}
