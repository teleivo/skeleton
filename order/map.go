package order

import (
	"cmp"
	"fmt"
	"io"
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

func (m Map[K, V]) RenderDot(w io.Writer) {
	fmt.Fprint(w, "strict digraph {\n")

	if m.IsEmpty() {
		fmt.Fprint(w, "}")
		return
	}

	if m.root.left == nil && m.root.right == nil {
		fmt.Fprintf(w, "\t%v\n}", m.root.key)
		return
	}

	visited := make(map[K]struct{})
	stack := [][2]*node[K, V]{
		{m.root, nil},
	}

	for len(stack) > 0 {
		x := stack[len(stack)-1][0]

		if x.left != nil && !isVisited(visited, x.left) {
			stack = append(stack, [2]*node[K, V]{x.left, x})
			continue
		}

		visited[x.key] = struct{}{}
		parent := stack[len(stack)-1][1]
		drawEdge(w, parent, x)

		stack = stack[:len(stack)-1]
		if x.right != nil {
			stack = append(stack, [2]*node[K, V]{x.right, x})
		}
	}

	fmt.Fprint(w, "}")
}

func drawEdge[K cmp.Ordered, V any](w io.Writer, from, to *node[K, V]) {
	if from == nil {
		return
	}

	label := "R"
	if from.left == to {
		label = "L"
	}
	var color string
	if to.red {
		color = ", color = red"
	}
	fmt.Fprintf(w, "\t%v -> %v [label=%q%s]\n", from.key, to.key, label, color)
}

func isRed[K cmp.Ordered, V any](n *node[K, V]) bool {
	if n == nil {
		return false
	}
	return n.red
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
