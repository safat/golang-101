package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func main() {
	root := Node{data: 1}

	node1 := Node{data: 2}
	root.next = &node1

	node2 := Node{data: 3}
	node1.next = &node2

	node3 := Node{data: -2}
	node2.next = &node3

	node4 := Node{data: -3}
	node3.next = &node4

	_ = node4

	Print2(&root)
	fmt.Println("----------------")
	Print(root)

}

func Print(head Node) {
	for head != (Node{}) {
		fmt.Println(head.data)

		if head.next != nil {
			head = *head.next
		} else {
			head = Node{}
		}
	}
}

func Print2(head *Node) {
	for head != nil {
		fmt.Println(head.data)

		head = head.next
	}
}

//https://www.hackerrank.com/challenges/print-the-elements-of-a-linked-list/problem