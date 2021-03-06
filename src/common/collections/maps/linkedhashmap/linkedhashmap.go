package linkedhashmap

import (
	"fmt"
	"strings"

	"github.com/sanksons/reflorest/src/common/collections"
)

// Link represents a node of doubly linked list
type Link struct {
	key   interface{}
	value interface{}
	next  *Link
	prev  *Link
}

// Map holds the elements in go's native map, also maintains the head and tail link
// to keep the elements in insertion order
type Map struct {
	m    map[interface{}]*Link
	head *Link
	tail *Link
}

func newLink(key interface{}, value interface{}) *Link {
	return &Link{key: key, value: value, next: nil, prev: nil}
}

// New instantiates a linked hash map.
func New() *Map {
	return &Map{m: make(map[interface{}]*Link), head: nil, tail: nil}
}

// Put inserts an element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	link, found := m.m[key]
	if !found {
		link = newLink(key, value)
		if m.tail == nil {
			m.head = link
			m.tail = link
		} else {
			m.tail.next = link
			link.prev = m.tail
			m.tail = link
		}
		m.m[key] = link
	} else {
		link.value = value
	}
}

// Get searches the element in the map by key and returns its value or nil if key doesn't exists.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	var link *Link
	link, found = m.m[key]
	if found {
		value = link.value
	} else {
		value = nil
	}
	return
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	link, found := m.m[key]
	if found {
		delete(m.m, key)
		if m.head == link && m.tail == link {
			m.head = nil
			m.tail = nil
		} else if m.tail == link {
			m.tail = link.prev
			link.prev.next = nil
		} else if m.head == link {
			m.head = link.next
			link.next.prev = nil
		} else {
			link.prev.next = link.next
			link.next.prev = link.prev
		}
	}
}

// IsEmpty returns true if map does not contain any elements
func (m *Map) IsEmpty() bool {
	return m.Size() == 0
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Keys returns all keys of the map (insertion order).
func (m *Map) Keys() []interface{} {
	keys := make([]interface{}, m.Size())
	count := 0
	for current := m.head; current != nil; current = current.next {
		keys[count] = current.key
		count++
	}
	return keys
}

// Values returns all values of the map (insertion order).
func (m *Map) Values() []interface{} {
	values := make([]interface{}, m.Size())
	count := 0
	for current := m.head; current != nil; current = current.next {
		values[count] = current.value
		count++
	}
	return values
}

// Contains returns true if the given keys are found in the map
func (m *Map) Contains(keys ...interface{}) bool {
	for _, key := range keys {
		_, found := m.m[key]
		if !found {
			return false
		}
	}
	return true
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.m = make(map[interface{}]*Link)
	m.head = nil
	m.tail = nil
}

// Iterator returns a stateful iterator used for iterating over all the entries of the map
func (m *Map) Iterator() collections.Iterator {
	return &Iterator{m: m, current: m.head}
}

// String returns a string representation of container
func (m *Map) String() string {
	str := "LinkedHashMap\nmap["
	for current := m.head; current != nil; current = current.next {
		str += fmt.Sprintf("%v:%v ", current.key, current.value)
	}
	return strings.TrimRight(str, " ") + "]"
}
