package main

import "code.google.com/p/go-tour/tree"

import (
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int, level int) {
	if t.Left != nil {
		Walk(t.Left, ch, level+1)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch, level+1)
	}
	if level == 0 {
		close(ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)
	go Walk(t1, c1, 0)
	go Walk(t2, c2, 0)
	for v1 := range c1 {
		v2, c2_open := <-c2
		if !c2_open || (v1 != v2) {
			return false
		}
	}
	// Check that c2 is also closed:
	_, c2_open := <-c2
	return true && !c2_open
}

func main() {
	c := make(chan int)
	go Walk(tree.New(1), c, 0)
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
