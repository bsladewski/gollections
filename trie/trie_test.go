package trie_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/bsladewski/gollections/trie"
)

// TestTrie tests all exported functionality of the Trie type.
func TestTrie(t *testing.T) {
	add := []string{"car", "cart", "cat", "three", "tree", "zebra"}
	addRemove := []string{"can", "care", "eat", "tame", "undo", "zen"}
	tr := trie.NewTrie()
	// complete, contains, remove work; empty trie
	if result := tr.Complete("test"); !reflect.DeepEqual([]string{}, result) {
		t.Fatalf("expected empty slice, got %v", result)
	}
	if tr.Contains("test") {
		t.Fatal("expected contains to return false on empty trie")
	}
	tr.Remove("test")
	// add
	tr.Add(add...)
	tr.Add(addRemove...)
	if !tr.Contains(add...) || !tr.Contains(addRemove...) {
		t.Error("expected contains to return true for all added elements")
	}
	expected := append(add, addRemove...)
	sort.Strings(expected)
	got := tr.Complete("")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// remove
	tr.Remove(addRemove...)
	if tr.Contains(addRemove...) {
		t.Error("expected contains to return false for removed elements")
	}
	if !tr.Contains(add...) {
		t.Error("expected elements not removed to still exist")
	}
	expected = add
	got = tr.Complete("")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// complete
	expected = []string{"car", "cart", "cat"}
	got = tr.Complete("ca")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
