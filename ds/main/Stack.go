package main

import "fmt"

type Element struct {
	value interface{}
	next  *Element
}

type Stack struct {
	size int
	top  *Element
}

func (stack *Stack) Push(value interface{}) {
	stack.top = &Element{value, stack.top}
	stack.size++
}

func (stack *Stack) Pop() (value interface{}) {
	if stack.size > 0 {
		value, stack.top = stack.top, stack.top.next
		stack.size--

		return
	}

	return nil
}

func (stack *Stack) IsEmpty() (bool) {
	return stack.size == 0
}

func main() {
	var stack = Stack{}

	stack.Push("java")
	stack.Push("php")
	stack.Push("c")
	stack.Push("cpp")
	stack.Push("go")

	fmt.Println(stack.Pop(), stack.size, stack.IsEmpty())
	fmt.Println(stack.Pop(), stack.size, stack.IsEmpty())
	fmt.Println(stack.Pop(), stack.size, stack.IsEmpty())
	fmt.Println(stack.Pop(), stack.size, stack.IsEmpty())
	fmt.Println(stack.Pop(), stack.size, stack.IsEmpty())

}
