package cache_test

import (
	"testing"

	"github.com/bsladewski/gollections"
	"github.com/bsladewski/gollections/cache"
)

// TestCache tests all exported functionality of a standard cache.
func TestCache(t *testing.T) {
	c := cache.NewCache(3)
	// get, size, remove; empty cache
	if got, err := c.Get("a"); got != nil || err != gollections.ErrNoSuchElement {
		t.Fatalf("expected nil and no such element error, got %v, err: %v", got, err)
	}
	if size := c.Size(); size != 0 {
		t.Fatalf("expected size 0, got %d", size)
	}
	c.Remove("a")
	// put, get, size
	for i := 0; i < 4; i++ {
		c.Put(i, 100-i)
	}
	if size := c.Size(); size != 3 {
		t.Fatalf("expected size 3, got %d", size)
	}
	for i := 3; i >= 0; i-- {
		got, err := c.Get(i)
		if i == 0 {
			if got != nil || err != gollections.ErrNoSuchElement {
				t.Fatalf("expected no such element error, got %v, err %v", got, err)
			}
			continue
		}
		if err != nil || got != 100-i {
			t.Fatalf("expected %d, got %d, err: %v", 100-i, got, err)
		}
	}
	// set max size
	c.SetMaxSize(2)
	if size := c.Size(); size != 2 {
		t.Fatalf("expected size 2, got %d", size)
	}
	for i := 0; i < 4; i++ {
		got, err := c.Get(i)
		if i == 0 || i == 3 {
			if got != nil || err != gollections.ErrNoSuchElement {
				t.Fatalf("expected no such element error, got %v, err %v", got, err)
			}
			continue
		}
		if err != nil || got != 100-i {
			t.Fatalf("expected %d, got %d, err: %v", 100-i, got, err)
		}
	}
	// remove
	c.Remove(1)
	if size := c.Size(); size != 1 {
		t.Fatalf("expected size 1, got %d", size)
	}
	if got, err := c.Get(1); got != nil || err != gollections.ErrNoSuchElement {
		t.Fatalf("expected no such element error, got %v, err %v", got, err)
	}
	// clear
	c.Clear()
	if size := c.Size(); size != 0 {
		t.Fatalf("expected size 0, got %d", size)
	}
	if got, err := c.Get(2); got != nil || err != gollections.ErrNoSuchElement {
		t.Fatalf("expected no such element error, got %v, err %v", got, err)
	}
}
