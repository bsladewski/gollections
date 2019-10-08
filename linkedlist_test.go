package gollections_test

import (
	"reflect"
	"testing"

	"github.com/bsladewski/gollections"
)

// Test the linkedList as an implementation of Collection.
func TestLinkedCollection(t *testing.T) {
	collection := gollections.NewLinkedCollection()
	// clear, contains, is empty, remove, size, to array; empty collection
	collection.Clear()
	if collection.Contains(0) {
		t.Fatal("expected contains to return false on empty collection")
	}
	if !collection.IsEmpty() {
		t.Fatal("expected is empty to return true for an empty collection")
	}
	collection.Remove(0)
	if size := collection.Size(); size != 0 {
		t.Fatalf("expected size 0 on empty collection, got %d", size)
	}
	if array := collection.ToArray(); len(array) != 0 {
		t.Fatalf("expected empty array, got %v", array)
	}
	// add
	expected := []interface{}{5, 3, 8, 4, 2, 6, 9}
	collection.Add(expected...)
	if size := collection.Size(); size != 7 {
		t.Errorf("expected size 7, got %d", size)
	}
	if got := collection.ToArray(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// remove
	collection.Remove(5, 9, 4, 0)
	if size := collection.Size(); size != 4 {
		t.Errorf("expected size 4, got %d", size)
	}
	expected = []interface{}{3, 8, 2, 6}
	if got := collection.ToArray(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// add; collection with tail removed
	expected = []interface{}{3, 8, 2, 6, 7}
	collection.Add(7)
	if got := collection.ToArray(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// clear
	collection.Clear()
	if size := collection.Size(); size != 0 {
		t.Errorf("expected size 0, got %d", size)
	}
	if got := collection.ToArray(); !reflect.DeepEqual([]interface{}{}, got) {
		t.Fatalf("expected empty collection, got %v", got)
	}
}

// TestLinkedCollectionSliceCopy tests the slice copy function for the linkedList implementation
// of Collection.
func TestLinkedCollectionSliceCopy(t *testing.T) {
	list := gollections.NewLinkedList()
	// slice copy; empty list
	got := &[]int{}
	if err := list.SliceCopy(got); err != nil {
		t.Fatalf("expected empty array, got %v", *got)
	}
	// slice copy; list with ints
	expected := []int{1, 2, 3}
	list.Add(1, 2, 3)
	if err := list.SliceCopy(got); err != nil || !reflect.DeepEqual(expected, *got) {
		t.Fatalf("expected %v, got %v", expected, *got)
	}
	// slice copy; list with a string
	list.Add("foo")
	if err := list.SliceCopy(got); err == nil {
		t.Fatalf("expected type error, got %v error: %v", *got, err)
	}
}

// Test the linkedList as an implementation of List.
func TestLinkedList(t *testing.T) {
	list := gollections.NewLinkedList()
	// index of, insert, get, remove at, set; empty list
	if index := list.IndexOf(0); index != -1 {
		t.Fatalf("expected -1, got %d", index)
	}
	if err := list.Insert(3, 6); err != gollections.ErrIndexOutOfBounds {
		t.Fatalf("expected index out of bounds error, got %v", err)
	}
	if _, err := list.Get(3); err != gollections.ErrIndexOutOfBounds {
		t.Fatalf("expected index out of bounds error, got %v", err)
	}
	if err := list.RemoveAt(3); err != gollections.ErrIndexOutOfBounds {
		t.Fatalf("expected index out of bounds error, got %v", err)
	}
	if err := list.Set(3, 6); err != gollections.ErrIndexOutOfBounds {
		t.Fatalf("expected index out of bounds error, got %v", err)
	}
	// index of, insert, get, remove at, set
	list.Add(6, 8, 3, 4, 5)
	if index := list.IndexOf(6); index != 0 {
		t.Fatalf("expected index 0, got %d", index)
	}
	if index := list.IndexOf(3); index != 2 {
		t.Fatalf("expected index 2, got %d", index)
	}
	if index := list.IndexOf(5); index != 4 {
		t.Fatalf("expected index 4, got %d", index)
	}
	expected := []interface{}{2, 7, 6, 8, 6, 3, 4, 9, 5}
	list.Insert(0, 2, 7)
	list.Insert(4, 6)
	list.Insert(7, 9)
	if got := list.ToArray(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	if got, err := list.Get(3); err != nil || got != 8 {
		t.Fatalf("expected 8, got %v, err: %v", got, err)
	}
	expected = []interface{}{7, 6, 8, 3, 4, 9}
	list.RemoveAt(0)
	list.RemoveAt(3)
	list.RemoveAt(6)
	if got := list.ToArray(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	expected = []interface{}{1, 6, 2, 3, 4, 7}
	list.Set(0, 1)
	list.Set(2, 2)
	list.Set(5, 7)
	if got := list.ToArray(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
}

// Test the linkedList as an implementation of Queue.
func TestLinkedQueue(t *testing.T) {
	queue := gollections.NewLinkedQueue()
	// peek first, pop first; empty queue
	if _, err := queue.PeekFirst(); err != gollections.ErrNoSuchElement {
		t.Fatalf("expected no such element error, got %v", err)
	}
	if _, err := queue.PopFirst(); err != gollections.ErrNoSuchElement {
		t.Fatalf("expected no such element error, got %v", err)
	}
	// peek first, pop first
	queue.Add(0, 1, 2, 3)
	for i := 0; i < 4; i++ {
		if got, err := queue.PeekFirst(); err != nil || i != got {
			t.Fatalf("expected %d, got %v", i, got)
		}
		if got, err := queue.PopFirst(); err != nil || i != got {
			t.Fatalf("expected %d, got %v", i, got)
		}
	}
	if !queue.IsEmpty() {
		t.Fatal("expected queue to be empty")
	}
}

// Test the linkedList as an implementation of Deque.
func TestLinkedDeque(t *testing.T) {
	deque := gollections.NewLinkedDeque()
	// peek last, pop last; empty deque
	if _, err := deque.PeekLast(); err != gollections.ErrNoSuchElement {
		t.Fatalf("expected no such element error, got %v", err)
	}
	if _, err := deque.PopLast(); err != gollections.ErrNoSuchElement {
		t.Fatalf("expected no such element error, got %v", err)
	}
	// add first
	deque.AddFirst(0, 1, 2, 3)
	expected := []interface{}{3, 2, 1, 0}
	if got := deque.ToArray(); !reflect.DeepEqual(expected, got) {
		t.Fatalf("expected %v, got %v", expected, got)
	}
	// peek last, pop last
	for i := 0; i < 4; i++ {
		if got, err := deque.PeekLast(); err != nil || i != got {
			t.Fatalf("expected %d, got %v", i, got)
		}
		if got, err := deque.PopLast(); err != nil || i != got {
			t.Fatalf("expected %d, got %v", i, got)
		}
	}
	if !deque.IsEmpty() {
		t.Fatal("expected deque to be empty")
	}
}
