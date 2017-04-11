package main

import "fmt"

// 用数组实现stack几乎是性能最优的方式
type StackBySlice []interface{}

func main() {
	stack := new(StackBySlice)
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

func (s StackBySlice) IsEmpty() bool {
	return len(s) == 0
}

func (s StackBySlice) Size() int {
	return len(s)
}

func (s *StackBySlice) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *StackBySlice) Pop() interface{} {
	n := len(*s) - 1
	result := (*s)[n]
	*s = (*s)[:n]
	return result
}
