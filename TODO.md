* add godoc badge

## Map

* implement LLRB tree
  * DeleteMin()
  * write an invariant assertion function I can run on a LLRB
  * implement Delete
  * should Min() also return the value like I do with DeleteMin? or should I not return the value
  from DeleteMin? With Min() you can Get() the value if needed, with DeleteMin() you could not
  * Keys(lo, hi key)
  * Backward()
  * could I support keys that implement a `func Less(i, j int) bool` as in https://pkg.go.dev/sort#Interface keys as well while not making it awkward for
  cmp.Ordered types? or maybe a https://pkg.go.dev/time#Time.Compare one as that also exists in
  cmp.Ordered. It would be cool to be able to use time.Time as the key to some value
  * how can I benchmark/test to make sure it is close to perfectly balanced?
  * Size() int
  * Size(lo, hi key)
  * Floor(key)
  * Ceiling(key)
  * Rank(key)
  * Select(k int)
  * Max() key
  * DeleteMax()
  * visualize tree - as in the lecture/paper I would like to visualize the tree on the CLI during
  sorted key insert and random order key insert; ideally with red links showing up as red ;)
* augment as described in lecture?

* testing
  * fix assertive - make a 0.0.1 release and use that, somehow updating does not work even though I
    did push a fix to the assertion message
  * look at coverage, did I miss any of the edge cases? rotations, moves?
  * can I write a fuzz test
  * how can I assert the invariants? how to assert in Go like for example in Java, then run it for
  example using a fuzz test so that I get the input that fails the invariant?
  * it would be cool to draw the expected tree, use that as the want value in the test assertion and
    show a visual diff if the test fails

* drawing
  * serialize the tree to a .dot file so I can debug the state at any time
    * invisible nodes are needed if I wanted the layout to show that the tree is indeed balanced
    https://forum.graphviz.org/t/how-to-get-trees-more-balanced/966/5

invisible nodes/edges with null links?
```go
`
strict digraph {
    10 -> 5 [color = red]
    5 -> 3
    3 -> 2
    2 -> "2L" [arrowhead=none]
    2 -> "2R" [arrowhead=none]
    "2L" [style=invis]
    "2R" [style=invis]
    3 -> 4
    5 -> 7
    7 -> 6
    7 -> 9
    9 -> 8 [color = red]
    9 -> "9R" [arrowhead=none]
    "9R" [style=invis]
    10 -> 20
    20 -> 15
    15 -> "15L" [arrowhead=none]
    15 -> "15R" [arrowhead=none]
    "15L" [style=invis]
    "15R" [style=invis]
    20 -> 23
    23 -> "23L" [arrowhead=none]
    23 -> "23R" [arrowhead=none]
    "23L" [style=invis]
    "23R" [style=invis]
}
`
```


## Future data structures

* make a priority queue using LLRB as described in lecture?
* make an ordered set: its basically the ordered table with a simpler API

