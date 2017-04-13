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

func judge(input [10]int, raw [10]int) bool {
	stack := new(Stack)
	inputTop := 0
	for i := 0; i < 10; i++ {
		stack.Push(raw[i])
		fmt.Println("push raw",raw[i])
		for stack.Top() == input[inputTop] {
			//fmt.Println("top and input", stack.Top(), input[inputTop])
			fmt.Println("Pop ", stack.Top())
			stack.Pop()
			inputTop++
			//fmt.Println("inputTop ", inputTop)
			if stack.Size() == 0 {
				break
			}
		}
	}
	//fmt.Println("end",inputTop)
	return inputTop == 10
}

func main() {
	inputa := [10]int{4, 3, 2, 1, 0, 9, 8, 7, 6, 5}
	inputb := [10]int{4, 6, 8, 7, 5, 3, 2, 9, 0, 1}
	inputc := [10]int{2, 5, 6, 7, 4, 8, 9, 3, 1, 0}
	inputd := [10]int{4, 3, 2, 1, 0, 5, 6, 7, 8, 9}
	inpute := [10]int{1, 2, 3, 4, 5, 6, 9, 8, 7, 0}
	inputf := [10]int{0, 4, 6, 5, 3, 8, 1, 7, 2, 9}
	inputg := [10]int{1, 4, 7, 9, 8, 6, 5, 3, 0, 2}
	inputh := [10]int{2, 1, 4, 3, 6, 5, 8, 7, 9, 0}
	raw := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("a",judge(inputa, raw))
	fmt.Println("b",judge(inputb, raw))
	fmt.Println("c",judge(inputc, raw))
	fmt.Println("d",judge(inputd, raw))
	fmt.Println("e",judge(inpute, raw))
	fmt.Println("f",judge(inputf, raw))
	fmt.Println("g",judge(inputg, raw))
	fmt.Println("h",judge(inputh, raw))
}
