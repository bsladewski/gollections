package gollections_test

import (
	"reflect"
	"sort"
	"testing"

	"github.com/bsladewski/gollections"
)

// TestTrie tests all exported functionality of the Trie type.
func TestTrie(t *testing.T) {
	// we will add and retain these strings in the trie
	add := []string{"car", "cart", "cat", "three", "tree", "zebra"}
	// we will add then later remove these strings from the trie
	addRemove := []string{"can", "care", "eat", "tame", "undo", "zen"}
	// create an empty trie
	trie := &gollections.Trie{}
	// ensure complete, contains, and remove work on an empty trie
	if result := trie.Complete("test"); !reflect.DeepEqual([]string{}, result) {
		t.Fatalf("expected empty slice, got %v", result)
	}
	if trie.Contains("test") {
		t.Fatal("expected contains to return false")
	}
	trie.Remove("test")
	// test adding elements to the trie
	trie.Add(add...)
	trie.Add(addRemove...)
	// check if contains indicates all added elements are present
	if !trie.Contains(add...) || !trie.Contains(addRemove...) {
		t.Error("expected contains to return true for all added elements")
	}
	// use complete to get all elements in the trie and compare added elements
	expected := append(add, addRemove...)
	sort.Strings(expected)
	got := trie.Complete("")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// test removing elements from the trie
	trie.Remove(addRemove...)
	// check if contains indicates removed elements are gone
	if trie.Contains(addRemove...) {
		t.Error("expected contains to return false for removed elements")
	}
	if !trie.Contains(add...) {
		t.Error("expected elements not removed to still exist")
	}
	// use complete to get all elements in the trie and compare to remaining elements
	expected = add
	got = trie.Complete("")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// test complete function on prefix "ca"
	expected = []string{"car", "cart", "cat"}
	got = trie.Complete("ca")
	sort.Strings(got)
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}
