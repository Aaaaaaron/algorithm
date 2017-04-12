package main

import (
	"fmt"
	"log"
)

type Stack []interface{}

func (s Stack) IsEmpty() bool {
	return len(s) == 0
}

func (s Stack) Size() int {
	return len(s)
}

func (s *Stack) Push(item interface{}) {
	*s = append(*s, item)
}

func (s *Stack) Pop() interface{} {
	n := len(*s) - 1
	result := (*s)[n]
	*s = (*s)[:n]
	return result
}

func (s *Stack) Top() interface{} {
	if s.Size() <= 0 {
		log.Fatal("null stack")
	}
	return (*s)[s.Size()-1 ]
}

func judge(input [10]int, raw [10]int) {
	stack := new(Stack)
	inputTop := 0
	//rawTop := 0
	for i := 0; i < 10; i++ {
		stack.Push(raw[i])
		fmt.Println("inputTop ", inputTop)
		for stack.Top() == input[inputTop] {
			fmt.Println("top and input", stack.Top(), input[inputTop])
			fmt.Println("Pop ", stack.Pop())
			inputTop++
		}
	}
	fmt.Println(inputTop)
}

func main() {
	input1 := [10]int{4, 3, 2, 1, 0, 9, 8, 7, 6, 5}
	//input2 := [10]int{4, 3, 2, 1, 0, 9, 8, 7, 6, 5}
	raw := [10]int{0, 1, 2, 3, 4, 5, 6, 6, 7, 8}
	judge(input1, raw)
	//stack := new(Stack)
	//stack.Push("a")
	//stack.Push("a")
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
}
