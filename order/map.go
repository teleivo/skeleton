package order

import (
	"cmp"
	"iter"
)

// Map is an ordered symbol table as described in the textbook Algorithms, 4th Edition by Robert
// Sedgewick and Kevin Wayne. A summary of the API can be found in
// https://algs4.cs.princeton.edu/31elementary. Map is implemented as a left-leaning red-black
// binary search tree as described in the paper Left-leaning Red-Black Trees by Robert Sedgewick
// https://sedgewick.io/wp-content/themes/sedgewick/papers/2008LLRB.pdf. All operations are thus
// guaranteed to run in O(log N) with N number of keys.
type Map[K cmp.Ordered, V any] struct {
	root *node[K, V]
}

type node[K cmp.Ordered, V any] struct {
	red         bool // Indicates if the link to this node is a red link
	key         K
	value       V
	left, right *node[K, V]
}

// IsEmpty returns true if the map contains no keys and false otherwise.
func (m *Map[K, V]) IsEmpty() bool {
	return m.root == nil
}

// Put associates the given key with the given value. The value of an existing key is updated.
// TODO allow nil value to delete the key?
func (m *Map[K, V]) Put(key K, value V) {
	m.root = m.put(m.root, key, value)
	m.root.red = false
}

func (m *Map[K, V]) put(n *node[K, V], key K, value V) *node[K, V] {
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
		n.left = m.put(n.left, key, value)
	} else {
		n.right = m.put(n.right, key, value)
	}

	return fixUp(n)
}

// Get returns the value associated with the given key and true if the key was found. The zero value
// and false is returned if the key was not found.
func (m *Map[K, V]) Get(key K) (V, bool) {
	for n := m.root; n != nil; {
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
func (m *Map[K, V]) Contains(key K) bool {
	_, ok := m.Get(key)
	return ok
}

func dfs[K cmp.Ordered, V any](n *node[K, V]) {
	if n == nil {
		return
	}

	if n.left != nil {
		dfs(n.left)
	}
	// n.key
	if n.right != nil {
		dfs(n.right)
	}
}

// All returns an in iterator iterating in ascending order over all key-value pairs.
func (m *Map[K, V]) All() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		visited := make(map[K]struct{})
		stack := []*node[K, V]{
			m.root,
		}

		for len(stack) > 0 {
			x := stack[len(stack)-1]

			if x.left != nil && !isVisited(visited, x.left) {
				stack = append(stack, x.left)
				continue
			}

			visited[x.key] = struct{}{}
			if !yield(x.key, x.value) {
				return
			}

			stack = stack[:len(stack)-1]
			if x.right != nil {
				stack = append(stack, x.right)
			}
		}
	}
}

// isVisited is needed to simplify to use the optimized 'set' with the empty struct value. Due to
// that optimization I cannot use visited in a boolean expression. This function allows that and
// combined with the x.left != nil guard simplifies the code and shrinks the number of cases I have
// to deal with.
func isVisited[K cmp.Ordered, V any](visited map[K]struct{}, x *node[K, V]) bool {
	_, ok := visited[x.key]
	return ok
}

// Min returns the smallest key in the map and true if the map is not empty. The zero value and
// false is returned if the map is empty.
func (m *Map[K, V]) Min() (K, bool) {
	var result K

	if m.root == nil {
		return result, false
	}

	for x := m.root; x != nil; {
		result = x.key
		x = x.left
	}

	return result, true
}

// DeleteMin returns the smallest key in the map, its associated value and true if the table is
// not empty. The zero value for the key and value and false is returned if the map is empty.
func (ta *Map[K, V]) DeleteMin() (K, V, bool) {
	if ta.IsEmpty() {
		var key K
		var value V
		return key, value, false
	}

	root, deleted := ta.deleteMin(ta.root)
	ta.root = root
	ta.root.red = false

	return deleted.key, deleted.value, true
}

func (ta *Map[K, V]) deleteMin(n *node[K, V]) (*node[K, V], *node[K, V]) {
	if n.left == nil {
		return nil, n
	}

	// TODO nil panic?
	if !isRed(n.left) && !isRed(n.left.left) {
		n = moveRedLeft(n)
	}

	left, deleted := ta.deleteMin(n.left)
	n.left = left

	return fixUp(n), deleted
}

func isRed[K cmp.Ordered, V any](n *node[K, V]) bool {
	if n == nil {
		return false
	}
	return n.red
}

func moveRedLeft[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
	flipColor(n)
	if isRed(n.right.left) {
		n.right = rotateRight(n)
		n = rotateLeft(n)
		flipColor(n)
	}
	return n
}

func fixUp[K cmp.Ordered, V any](n *node[K, V]) *node[K, V] {
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
