package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func InOrderTraverse(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	InOrderTraverse(t.Left, ch)
	ch <- t.Value
	InOrderTraverse(t.Right, ch)
}

func main() {
	tree1 := tree.New(10)
	tree2 := tree.New(10)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		InOrderTraverse(tree1, ch1)
	}()

	go func() {
		defer close(ch2)
		InOrderTraverse(tree2, ch2)
	}()

	var equalTree = true

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if ok1 != ok2 {
			equalTree = false
			break
		}

		if v1 != v2 {
			equalTree = false
			break
		}

		if ok1 == false && ok2 == false {
			break
		}
	}

	fmt.Println("Is Equal Tree: ", equalTree)

}
