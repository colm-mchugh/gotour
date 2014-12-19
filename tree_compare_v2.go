package main

import (
	"code.google.com/p/go-tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

func do_tree_walk(t *tree.Tree) chan int {
	c := make(chan int)
	go func() {
		Walk(t, c)
		close(c)
	}()
	return c
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := do_tree_walk(t1)
	c2 := do_tree_walk(t2)
	for {
		v1, c1_open := <-c1
		v2, c2_open := <-c2
		if !c1_open || !c2_open {
			return c1_open == c2_open
		}
		if v1 != v2 {
			return false
		}
	}
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
