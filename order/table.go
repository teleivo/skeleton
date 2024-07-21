package order

// Table is an ordered symbol Table as described in the textbook Algorithms, 4th Edition by Robert
// Sedgewick and Kevin Wayne. A summary of the API can be found in
// https://algs4.cs.princeton.edu/31elementary. Table is implemented as a left-leaning red-black
// binary search tree as described in the paper Left-leaning Red-Black Trees by Robert Sedgewick
// https://sedgewick.io/wp-content/themes/sedgewick/papers/2008LLRB.pdf. All operations are thus
// guaranteed to run in O(log N) with N number of keys.
type Table struct {
	root *node
}

type node struct {
	red         bool
	key         int
	value       int
	left, right *node
}

// Put associates the given key with the given value.
// TODO allow nil value to delete the key?
func (ta *Table) Put(key, value int) {
	ta.root = ta.put(ta.root, key, value)
	ta.root.red = false
}

func (ta *Table) put(n *node, key, value int) *node {
	if n == nil {
		return &node{
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

func (ta *Table) Get(key int) (int, bool) {
	for n := ta.root; n != nil; {
		if n.key == key {
			return n.value, true
		} else if n.key > key {
			n = n.left
		} else {
			n = n.right
		}
	}

	return 0, false
}

func (ta *Table) Contains(key int) bool {
	_, ok := ta.Get(key)
	return ok
}

func isRed(n *node) bool {
	if n == nil {
		return false
	}
	return n.red
}

func rotateLeft(n *node) *node {
	x := n.right
	n.right = x.left
	x.left = n
	x.red = n.red
	n.red = true
	return x
}

func rotateRight(n *node) *node {
	x := n.left
	n.left = x.right
	x.right = n
	x.red = n.red
	n.red = true
	return x
}

func flipColor(n *node) {
	n.red = !n.red
	n.left.red = !n.left.red
	n.right.red = !n.right.red
}
