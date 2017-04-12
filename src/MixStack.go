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
		fmt.Println("push raw",raw[i])
		for stack.Top() == input[inputTop] {
			fmt.Println("top and input", stack.Top(), input[inputTop])
			fmt.Println("Pop ", stack.Pop())
			inputTop++
			fmt.Println("inputTop ", inputTop)
			if stack.Size() == 0 {
				break
			}
		}
	}
	fmt.Println("end",inputTop)
}

func main() {
	//input := [10]int{4, 3, 2, 1, 0, 9, 8, 7, 6, 5}
	//input := [10]int{4, 6, 8, 7, 5, 3, 2, 9, 0, 1}
	//input := [10]int{2, 5, 6, 7, 4, 8, 9, 3, 1, 0}
	input := [10]int{2, 1, 4, 3, 6, 5, 8, 7, 9, 0}
	raw := [10]int{0, 1, 2, 3, 4, 5, 6,7, 8, 9}
	judge(input, raw)
	//stack := new(Stack)
	//stack.Push("a")
	//stack.Push("a")
	//fmt.Println(stack.Pop())
	//fmt.Println(stack.Pop())
}
