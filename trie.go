package gollections

// A Trie is used to quickly check for and retrieve strings.
type Trie struct {
	value    rune
	children map[rune]*Trie
}

// adds the supplied string to the trie character by character.
func (trie *Trie) add(value []rune, index int) {
	if index >= len(value) {
		trie.children[0] = &Trie{}
		return
	}
	current := value[index]
	node, ok := trie.children[current]
	if !ok {
		node = &Trie{value: current, children: map[rune]*Trie{}}
		trie.children[current] = node
	}
	node.add(value, index+1)
}

// Add inserts new values into the trie.
func (trie *Trie) Add(values ...string) {
	if trie.children == nil {
		trie.children = map[rune]*Trie{}
	}
	for _, value := range values {
		trie.add([]rune(value), 0)
	}
}

// get finds a node in the trie using the supplied value as a path.
// Returns nil if no such node is found.
func (trie *Trie) get(value []rune, index int) *Trie {
	if index == len(value) {
		return trie
	}
	node, ok := trie.children[value[index]]
	if !ok {
		return nil
	}
	return node.get(value, index+1)
}

// traverse returns all strings that begin with the specified prefix.
// If no relevant strings exist, the resulting array will be empty.
func (trie *Trie) traverse(prefix string, first bool) []string {
	if !first && trie.value == 0 {
		return []string{prefix}
	}
	values := []string{}
	if !first {
		prefix += string(trie.value)
	}
	for _, v := range trie.children {
		values = append(values, v.traverse(prefix, false)...)
	}
	return values
}

// Complete returns all strings that complete the supplied string.
// If no relevant strings exist, the resulting array will be empty.
func (trie *Trie) Complete(value string) []string {
	node := trie.get([]rune(value), 0)
	if node == nil {
		return []string{}
	}
	return node.traverse(value, true)
}

// contains checks if the trie contains the specified value.
func (trie *Trie) contains(value []rune, index int) bool {
	if index == len(value) {
		_, ok := trie.children[0]
		return ok
	}
	node, ok := trie.children[value[index]]
	if !ok {
		return false
	}
	return node.contains(value, index+1)
}

// Contains checks if the trie contains all specified values.
func (trie *Trie) Contains(values ...string) bool {
	for _, value := range values {
		if !trie.contains([]rune(value), 0) {
			return false
		}
	}
	return true
}

// remove deletes the specified value from the trie.
// Returns true if the node should be removed from its parent.
func (trie *Trie) remove(value []rune, index int) bool {
	if index == len(value) {
		delete(trie.children, 0)
		return len(trie.children) == 0
	}
	node, ok := trie.children[value[index]]
	if !ok {
		return false
	}
	if node.remove(value, index+1) {
		delete(trie.children, value[index])
	}
	return len(trie.children) == 0
}

// Remove deletes the specified values from the trie.
func (trie *Trie) Remove(values ...string) {
	for _, value := range values {
		trie.remove([]rune(value), 0)
	}
}
