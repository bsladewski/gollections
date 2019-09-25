package gollections_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/bsladewski/gollections"
)

// TestTrie tests all exported functionality of the Trie type.
func TestTrie(t *testing.T) {
	add := []string{"car", "cart", "cat", "three", "tree", "zebra"}
	addRemove := []string{"can", "care", "eat", "tame", "undo", "zen"}
	trie := &gollections.Trie{}
	// complete, contains, remove work; empty trie
	if result := trie.Complete("test"); !reflect.DeepEqual([]string{}, result) {
		t.Fatalf("expected empty slice, got %v", result)
	}
	if trie.Contains("test") {
		t.Fatal("expected contains to return false on empty trie")
	}
	trie.Remove("test")
	// add
	trie.Add(add...)
	trie.Add(addRemove...)
	if !trie.Contains(add...) || !trie.Contains(addRemove...) {
		t.Error("expected contains to return true for all added elements")
	}
	expected := append(add, addRemove...)
	sort.Strings(expected)
	got := trie.Complete("")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// remove
	trie.Remove(addRemove...)
	if trie.Contains(addRemove...) {
		t.Error("expected contains to return false for removed elements")
	}
	if !trie.Contains(add...) {
		t.Error("expected elements not removed to still exist")
	}
	expected = add
	got = trie.Complete("")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// complete
	expected = []string{"car", "cart", "cat"}
	got = trie.Complete("ca")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
