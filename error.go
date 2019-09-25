package gollections

import "errors"

var (
	// ErrIndexOutOfBounds the supplied index was invalid for this list.
	ErrIndexOutOfBounds = errors.New("index out of bounds")

	// ErrNoSuchElement the polled element does not exist.
	ErrNoSuchElement = errors.New("no such element")
)
