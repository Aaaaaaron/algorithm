package main

import "fmt"

/*
0-9顺序入栈,以下哪种出栈序列是不可能的
*/

func judgePopOrder(input [10]int, raw [10]int) bool {
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
	fmt.Println("a", judgePopOrder(inputa, raw))
	fmt.Println("b", judgePopOrder(inputb, raw))
	fmt.Println("c", judgePopOrder(inputc, raw))
	fmt.Println("d", judgePopOrder(inputd, raw))
	fmt.Println("e", judgePopOrder(inpute, raw))
	fmt.Println("f", judgePopOrder(inputf, raw))
	fmt.Println("g", judgePopOrder(inputg, raw))
	fmt.Println("h", judgePopOrder(inputh, raw))
}
