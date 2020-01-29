package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch) // <- closes the channel when this function returns
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 && !ok2 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println("Equivalent Binary Trees\n")

	t1 := tree.New(5)
	t2 := tree.New(5)
	fmt.Println("Tree 1:", t1)
	fmt.Println("Tree 2:", t2)
	fmt.Println("Are they same? - ", Same(t1, t2))
	fmt.Println("------")

	t3 := tree.New(2)
	t4 := tree.New(3)
	fmt.Println("Tree 3:", t3)
	fmt.Println("Tree 4:", t4)
	fmt.Println("Are they same? - ", Same(t3, t4))
	fmt.Println("")
}
