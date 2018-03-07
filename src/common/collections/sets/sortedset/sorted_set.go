package sortedset

import (
	"github.com/sanksons/reflorest/src/common/collections"
	"github.com/sanksons/reflorest/src/common/collections/sets"
)

// Package SortedSet provides a Set that further provides a total ordering on its elements.
//
// The elements are ordered using their natural ordering, or by a
// Comparator typically provided at sorted set creation time.
type SortedSet interface {
	// First returns the first(min) entry in the set
	First() interface{}
	// Last returns the last(max) element in the set
	Last() interface{}
	// extends Set, Iterable and Comparable interfaces
	sets.Set
	collections.Iterable
	collections.Comparable
}
