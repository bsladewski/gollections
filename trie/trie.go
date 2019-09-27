// Package trie provides an implemenation of a trie data structure.
package trie

// A Trie is a set that is optimized for working with strings.
type Trie interface {
	// Add inserts new values into the trie.
	Add(values ...string)
	// Complete returns all strings that complete the supplied prefix string.
	// If no relevant strings exist, the resulting array will be empty.
	Complete(prefix string) []string
	// Contains checks if the trie contains all specified values.
	Contains(values ...string) bool
	// Remove deletes the specified values from the trie.
	Remove(values ...string)
}

// A trie is used to quickly check for and retrieve strings.
type trie struct {
	value    rune
	children map[rune]*trie
}

// adds the supplied string to the trie character by character.
func (t *trie) add(value []rune, index int) {
	if index >= len(value) {
		t.children[0] = &trie{}
		return
	}
	current := value[index]
	node, ok := t.children[current]
	if !ok {
		node = &trie{value: current, children: map[rune]*trie{}}
		t.children[current] = node
	}
	node.add(value, index+1)
}

func (t *trie) Add(values ...string) {
	if t.children == nil {
		t.children = map[rune]*trie{}
	}
	for _, value := range values {
		t.add([]rune(value), 0)
	}
}

// get finds a node in the trie using the supplied value as a path.
// Returns nil if no such node is found.
func (t *trie) get(value []rune, index int) *trie {
	if index == len(value) {
		return t
	}
	node, ok := t.children[value[index]]
	if !ok {
		return nil
	}
	return node.get(value, index+1)
}

// traverse returns all strings that begin with the specified prefix.
// If no relevant strings exist, the resulting array will be empty.
func (t *trie) traverse(prefix string, first bool) []string {
	if !first && t.value == 0 {
		return []string{prefix}
	}
	values := []string{}
	if !first {
		prefix += string(t.value)
	}
	for _, v := range t.children {
		values = append(values, v.traverse(prefix, false)...)
	}
	return values
}

func (t *trie) Complete(prefix string) []string {
	node := t.get([]rune(prefix), 0)
	if node == nil {
		return []string{}
	}
	return node.traverse(prefix, true)
}

// contains checks if the trie contains the specified value.
func (t *trie) contains(value []rune, index int) bool {
	if index == len(value) {
		_, ok := t.children[0]
		return ok
	}
	node, ok := t.children[value[index]]
	if !ok {
		return false
	}
	return node.contains(value, index+1)
}

func (t *trie) Contains(values ...string) bool {
	for _, value := range values {
		if !t.contains([]rune(value), 0) {
			return false
		}
	}
	return true
}

// remove deletes the specified value from the trie.
// Returns true if the node should be removed from its parent.
func (t *trie) remove(value []rune, index int) bool {
	if index == len(value) {
		delete(t.children, 0)
		return len(t.children) == 0
	}
	node, ok := t.children[value[index]]
	if !ok {
		return false
	}
	if node.remove(value, index+1) {
		delete(t.children, value[index])
	}
	return len(t.children) == 0
}

func (t *trie) Remove(values ...string) {
	for _, value := range values {
		t.remove([]rune(value), 0)
	}
}

// NewTrie initializes a new trie.
func NewTrie() Trie {
	return &trie{}
}
