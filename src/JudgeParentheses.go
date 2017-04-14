package main

import "fmt"

func main() {
	p1 := []string{"[", "(", ")", "]", "{", "}", "{", "[", "(", ")", "]", "}"}
	fmt.Println(judgeParentheses(p1))
	p2 := []string{"[", "(", "]", ")"}
	fmt.Println(judgeParentheses(p2))
}

func judgeParentheses(parentheses []string) bool {
	left := new(Stack)
	for _, value := range parentheses {
		if leftParentheses(value) {
			//fmt.Println("push", value)
			left.Push(value)
		} else {
			if !pairing(left.Top().(string), value) {
				return false
			}
			left.Pop()
			//fmt.Println("pop", value)
		}
	}
	return true
}

func pairing(l string, r string) bool {
	if (l == "(" && r == ")") || (l == "[" && r == "]") || (l == "{" && r == "}") {
		return true
	}
	return false
}

func leftParentheses(p string) bool {
	if p == "(" || p == "[" || p == "{" {
		return true
	} else {
		return false
	}
}
