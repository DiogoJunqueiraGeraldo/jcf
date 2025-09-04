package jcf

import "iter"

// Coll is the root interface in the collection hierarchy.
// A collection represents a group of objects, known as its elements.
// Some collections allow duplicate elements and others do not.
// Some are ordered and others unordered.
//
// https://docs.oracle.com/javase/8/docs/api/java/util/Collection.html
type Coll[T any] interface {
	// Contains returns true if this collection contains the specified element.
	Contains(elem T) bool

	// ContainsAll returns true if this collection contains all the elements in the specified collection.
	ContainsAll(coll Coll[T]) bool

	// Equals compares the specified object with this collection for equality.
	Equals(coll Coll[T]) bool

	// IsEmpty returns true if this collection contains no elements.
	IsEmpty() bool

	// Size returns the number of elements in this collection.
	Size() int

	// ToSlice returns a slice containing all the elements in this collection.
	ToSlice() []T

	// Copy returns an array containing all the elements in this collection;
	// the runtime type of the returned array is that of the specified array.
	Copy(out []T)
}

// ExtensibleColl is a collection that supports adding elements.
type ExtensibleColl[T comparable] interface {
	Coll[T]

	// Add ensures that this collection contains the specified element.
	Add(elem T) bool

	// AddAll adds all the elements in the specified collection to this collection.
	AddAll(coll Coll[T]) bool
}

// RemovableColl is a collection that supports removing elements.
type RemovableColl[T comparable] interface {
	Coll[T]

	// Remove removes a single instance of the specified element from this collection, if it is present.
	Remove(elem T) bool

	// RemoveAll removes all of this collection's elements that are also contained in the specified collection.
	RemoveAll(coll Coll[T]) bool

	// RemoveIf removes all the elements of this collection that satisfy the given predicate.
	RemoveIf(func(elem T) bool) bool

	// RetainAll retains only the elements in this collection that are contained in the specified collection.
	RetainAll(coll Coll[T]) bool

	// Clear removes all the elements from this collection.
	Clear()
}

// MutableColl is a collection that supports both adding and deleting elements
type MutableColl[T any] interface {
	ExtensibleColl[T]
	RemovableColl[T]
}

// IterableColl is a collection that supports iterating over elements.
type IterableColl[T any] interface {
	Coll[T]

	// Iter returns an iterator over the elements in this collection.
	Iter() iter.Seq[T]

	// Ch returns a channel that will produce all collection elements and close after draining
	Ch(buffSize int) <-chan T

	// ChParallel returns a channel that will produce all collection elements and close after draining doesn't conserve ordering
	ChParallel(buffSize, workers int) <-chan T
}
