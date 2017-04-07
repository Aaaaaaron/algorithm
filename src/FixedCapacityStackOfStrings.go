package main

import "fmt"

type StackByArray []interface{}

func main() {
	stack := new(StackByArray)
	fmt.Println(stack.IsEmpty())
	stack.Push("a")
	stack.Push("a")
	stack.Push("a")
	//stack.Push("b")
	//stack.Push("c")
	//stack.Push("d")
	//stack.Push("e")
	fmt.Println(cap(*stack))
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Size())
}

func (s StackByArray) IsEmpty() bool {
	return len(s) == 0
}

func (s StackByArray) Size() int {
	return len(s)
}

func (s *StackByArray) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *StackByArray) Pop() interface{} {
	n := len(*s) - 1
	result := (*s)[n]
	*s = (*s)[:n]
	return result
}
