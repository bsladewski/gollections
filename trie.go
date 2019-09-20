package gollections

// A Trie is used to quickly check for and retrieve strings.
type Trie struct {
	value    byte
	children map[string]*Trie
	length   int
}

// Add appends new elements to the collection.
func (trie *Trie) Add(values ...interface{}) {

}

// Clear removes all elements from the collection.
func (trie *Trie) Clear() {

}

// Contains checks if the collection contains all specified values.
func (trie *Trie) Contains(values ...interface{}) bool {
	return false
}

// IsEmpty checks if the collection contains no elements.
func (trie *Trie) IsEmpty() bool {
	return false
}

// Remove removes all specified values from the collection.
func (trie *Trie) Remove(values ...interface{}) {

}

// Size gets the number of elements in the collection.
func (trie *Trie) Size() int {
	return 0
}

// ToArray gets an array representation of the collection.
func (trie *Trie) ToArray() []interface{} {
	return nil
}
