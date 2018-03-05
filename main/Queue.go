package main

import "fmt"

type Item struct {
	value interface{}
	next  *Item
}

type Queue struct {
	size int
	top  *Item
}

func (queue *Queue) Push(value interface{}) {
	queue.size++

	item := Item{value: value}

	if queue.top == nil {
		queue.top = &item
		return
	}

	var lastItem = queue.top

	for lastItem.next != nil {
		lastItem = lastItem.next
	}

	lastItem.next = &item
}

func (queue *Queue) Pop() (value interface{}) {
	if queue.size == 0 {
		return nil
	}

	value = queue.top
	queue.top = queue.top.next
	queue.size--

	return value
}

func (queue *Queue) IsEmpty() (bool) {
	return queue.size == 0
}

func main() {
	var queue = Queue{}

	queue.Push("C")
	queue.Push("Java")
	queue.Push("F#")
	queue.Push("Angular")

	fmt.Println(queue.Pop(), queue.size, queue.IsEmpty())
	fmt.Println(queue.Pop(), queue.size, queue.IsEmpty())
	fmt.Println(queue.Pop(), queue.size, queue.IsEmpty())

	queue.Push("GOLANG")

	fmt.Println(queue.Pop(), queue.size, queue.IsEmpty())
	fmt.Println(queue.Pop(), queue.size, queue.IsEmpty())
}
