package main

import (
	"fmt"
	"strings"
)

func main() {
	expr := strings.Split("1 + 2 ) * 3 - 4 ) * 5 - 6 ) ) )", " ")
	fmt.Println(completeRightBrackets(expr))
}

func completeRightBrackets(expr []string) string {
	dataStack := new(Stack)
	opStack := new(Stack)

	for _, value := range expr {
		if value == ")" {
			exp2 := dataStack.Pop().(string)
			exp1 := dataStack.Pop().(string)
			dataStack.Push("( " + exp1 + " " + opStack.Pop().(string) + " " + exp2 + " )")
		} else if isOp(value) {
			opStack.Push(value)
		} else {
			dataStack.Push(value)
		}
	}
	return dataStack.Pop().(string)
}

func isOp(v string) bool {
	return v == "+" || v == "-" || v == "*" || v == "/"
}
