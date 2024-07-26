* add godoc badge

## Table

* implement LLRB tree
  * DeleteMin()
  * implement delete
  * should Min() also return the value like I do with DeleteMin? or should I not return the value
  from DeleteMin?
  * could I support keys that implement a `func Less(i, j int) bool` as in https://pkg.go.dev/sort#Interface keys as well while not making it awkward for
  cmp.Ordered types? or maybe a https://pkg.go.dev/time#Time.Compare one as that also exists in
  cmp.Ordered. It would be cool to be able to use time.Time as the key to some value
  * how can I benchmark/test to make sure it is close to perfectly balanced?
  * size() int
  * max() key
  * floor(key)
  * ceiling(key)
  * rank(key)
  * select(k int)
  * deleteMax()
  * keys(lo, hi key) - using Go 1.23 iterator?
  * keys()
  * finish frequency counter example - document the examples in the readme
  * size(lo, hi key)
  * visualize tree - as in the lecture/paper I would like to visualize the tree on the CLI during
  sorted key insert and random order key insert; ideally with red links showing up as red ;)
* augment as described in lecture?

* testing
  * can I write a fuzz test
  * how can I assert the invariants? how to assert in Go like for example in Java, then run it for
  example using a fuzz test so that I get the input that fails the invariant?
  * it would be cool to draw the expected tree, use that as the want value in the test assertion and
    show a visual diff if the test fails

## Future data structures

* make a priority queue using LLRB as described in lecture?
* make an ordered set: its basically the ordered table with a simpler API

