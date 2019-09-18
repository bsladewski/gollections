package gollections

// ListNode represents a single element in a doubly linked list.
type ListNode struct {
	Value    interface{}
	Previous *ListNode
	Next     *ListNode
}

// LinkedList is an implementation of a doubly linked list.
type LinkedList struct {
	Head   *ListNode
	Tail   *ListNode
	length int
}

// Add appends new elements to the end of the collection.
func (list *LinkedList) Add(values ...interface{}) {

}

// Clear removes all elements from the collection.
func (list *LinkedList) Clear() {

}

// Contains checks if the collection contains all specified values.
func (list *LinkedList) Contains(values ...interface{}) bool {
	return false
}

// IsEmpty checks if the collection contains no elements.
func (list *LinkedList) IsEmpty() bool {
	return false
}

// Remove removes all specified values from the collection.
func (list *LinkedList) Remove(value interface{}) {

}

// Size gets the number of elements in the collection.
func (list *LinkedList) Size() int {
	return 0
}

// ToArray gets an array representation of the collection.
func (list *LinkedList) ToArray() []interface{} {
	return nil
}

// IndexOf gets the first occurance of the specified value or -1 if not found.
func (list *LinkedList) IndexOf(value interface{}) int {
	return 0
}

// Insert adds elements at the specified index. Can return index not found error.
func (list *LinkedList) Insert(index int, values ...interface{}) error {
	return nil
}

// RemoveAt removes the element at the specified index.
func (list *LinkedList) RemoveAt(index int) error {
	return nil
}

// Set overwrites the value of the element at the specified index.
func (list *LinkedList) Set(index int, value interface{}) error {
	return nil
}

// PeekFirst gets the value of the first element in the collection.
func (list *LinkedList) PeekFirst() (interface{}, error) {
	return nil, nil
}

// PopFirst gets the value of the first element in the collection. The element is removed.
func (list *LinkedList) PopFirst() (interface{}, error) {
	return nil, nil
}

// PeekLast gets the value of the last element in the collection.
func (list *LinkedList) PeekLast() (interface{}, error) {
	return nil, nil
}

// PopLast gets the value of the last element in the collection. The element is removed.
func (list *LinkedList) PopLast() (interface{}, error) {
	return nil, nil
}

// AddFirst adds new elements to the beginning of the collection.
func (list *LinkedList) AddFirst(values ...interface{}) {

}
