package main

import (
	"fmt"
	"strings"
)

/*
把中缀表达式变为后缀,后缀的好处是没有括号(应该是没有优先级了,严格从左到右) 此处只考虑+-
*/

func main() {
	expr := strings.Split("9 + ( 3 - 1 ) * 3 + 10 / 2", " ")
	fmt.Println(convert(expr))
}

func convert(expr []string) string {
	opStack := new(Stack)
	postfixExpr := ""
	for _, value := range expr {
		if isOp2(value) {
			if opStack.Size() == 0 {
				opStack.Push(value)
				continue
			}

			if value == "(" {
				opStack.Push(value)
				fmt.Println("push" + value)
			}

			if value == ")" {
				for {
					if opStack.Top() == "(" {
						opStack.Pop()
						fmt.Println("pop" + "(")
						break
					}
					postfixExpr += opStack.Pop().(string)
					fmt.Println(") ", postfixExpr)
				}
			}

			if opStack.Size() > 0 && ( opPriority(value) > opPriority(opStack.Top().(string))) {
				opStack.Push(value)
				fmt.Println("push" + value)
			} else {
				for opPriority(value) <= opPriority(opStack.Top().(string)) {
					postfixExpr += opStack.Pop().(string)
					fmt.Println("pop ", postfixExpr)
					if opStack.IsEmpty() {
						break
					}
				}
				opStack.Push(value)
				fmt.Println("push" + value)
			}
		}
		postfixExpr += value
		fmt.Println("end ", postfixExpr)
	}
	return postfixExpr
}

func opPriority(value string) int {
	switch value {
	case "*", "/":
		return 2
	case "+", "-":
		return 1
	}
	return -1
}

func isOp2(v string) bool {
	return v == "+" || v == "-" || v == "*" || v == "/" || v == "(" || v == ")"
}
