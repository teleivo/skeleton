* implement LLRB tree
  * implement delete
  * min() key - also return the value?
  * deleteMin()
  * implement frequency counter example using data from lecture
  * could I support keys that implement a `func Less(i, j int) bool` as in https://pkg.go.dev/sort#Interface keys as well while not making it awkward for
  cmp.Ordered types? or maybe a https://pkg.go.dev/time#Time.Compare one as that also exists in
  cmp.Ordered. It would be cool to be able to use time.Time as the key to some value
  * how can I benchmark/test to make sure it is close to perfectly balanced?
  * isEmpty() bool
  * size() int
  * max() key
  * floor(key)
  * ceiling(key)
  * rank(key)
  * select(k int)
  * deleteMax()
  * size(lo, hi key)
  * keys(lo, hi key) - using Go 1.23 iterator?
  * keys()
  * visualize tree - as in the lecture/paper I would like to visualize the tree on the CLI during
  sorted key insert and random order key insert; ideally with red links showing up as red ;)
* augment as described in lecture?
* make a priority queue using it as described in lecture?
* testing
  * it would be cool to draw the expected tree, use that as the want value in the test assertion and
    show a visual diff if the test fails
