package order

import (
	"cmp"
)

// Table is an ordered symbol Table as described in the textbook Algorithms, 4th Edition by Robert
// Sedgewick and Kevin Wayne. A summary of the API can be found in
// https://algs4.cs.princeton.edu/31elementary. Table is implemented as a left-leaning red-black
// binary search tree as described in the paper Left-leaning Red-Black Trees by Robert Sedgewick
// https://sedgewick.io/wp-content/themes/sedgewick/papers/2008LLRB.pdf. All operations are thus
// guaranteed to run in O(log N) with N number of keys.
type Table[K cmp.Ordered, V any] struct {
	root *node[K, V]
}

type node[K cmp.Ordered, V any] struct {
	red         bool // Indicates if the link to this node is a red link
	key         K
	value       V
	left, right *node[K, V]
}

// Put associates the given key with the given value. The value of an existing key is updated.
// TODO allow nil value to delete the key?
func (ta *Table[K, V]) Put(key K, value V) {
	ta.root = ta.put(ta.root, key, value)
	ta.root.red = false
}

func (ta *Table[K, V]) put(n *node[K, V], key K, value V) *node[K, V] {
	if n == nil {
		return &node[K, V]{
			red:   true,
			key:   key,
			value: value,
		}
	}

	if n.key == key {
		n.value = value
	} else if n.key > key {
		n.left = ta.put(n.left, key, value)
	} else {
		n.right = ta.put(n.right, key, value)
	}

	if !isRed(n.left) && isRed(n.right) {
		n = rotateLeft(n)
	}
	if isRed(n.left) && isRed(n.left.left) {
		n = rotateRight(n)
	}
	if isRed(n.left) && isRed(n.right) {
		flipColor(n)
	}

	return n
}

// Get returns the value associated with the given key and true if the key was found. The zero value
// and false is returned if the key was not found.
func (ta *Table[K, V]) Get(key K) (V, bool) {
	for n := ta.root; n != nil; {
		if n.key == key {
			return n.value, true
		} else if n.key > key {
			n = n.left
		} else {
			n = n.right
		}
	}

	var result V
	return result, false
}

// Contains returns true if the given key was found and false otherwise.
func (ta *Table[K, V]) Contains(key K) bool {
	_, ok := ta.Get(key)
	return ok
}

// Min returns the smallest key in the table and true if the table is not empty. The zero value
// and false is returned if the table is empty.
func (ta *Table[K, V]) Min() (K, bool) {
	var result K

	if ta.root == nil {
		return result, false
	}

	for x := ta.root; x != nil; {
		result = x.key
		x = x.left
	}

	return result, true
}

func isRed[K cmp.Ordered, V any](n *node[K, V]) bool {
	if n == nil {
		return false
	}
	return n.red
}

func rotateLeft[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	x := n.right
	n.right = x.left
	x.left = n
	x.red = n.red
	n.red = true
	return x
}

func rotateRight[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	x := n.left
	n.left = x.right
	x.right = n
	x.red = n.red
	n.red = true
	return x
}

func flipColor[K cmp.Ordered, V any](n *node[K, V]) {
	n.red = !n.red
	n.left.red = !n.left.red
	n.right.red = !n.right.red
}
