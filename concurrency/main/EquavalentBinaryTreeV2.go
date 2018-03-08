package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

func inOrderTraverse(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	inOrderTraverse(t.Left, ch)
	ch <- t.Value
	inOrderTraverse(t.Right, ch)
}

func inOrderTraverseUtil(tree *tree.Tree, ch chan int) {
	defer close(ch)

	inOrderTraverse(tree, ch)
}

func main() {
	tree1 := tree.New(10)
	tree2 := tree.New(10)

	ch1 := make(chan int)
	ch2 := make(chan int)

	go inOrderTraverseUtil(tree1, ch1)
	go inOrderTraverseUtil(tree2, ch2)

	equalTree := true

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

		if !ok1 && !ok2 {
			break
		}
	}

	fmt.Println("Is Equal Tree: ", equalTree)
}
