package main

import "fmt"

var in Stack
var out Stack

func push(s string) {
	in.Push(s)
}

func pop() string {
	if out.Size() == 0 {
		for in.Size() > 0 {
			out.Push(in.Pop())
		}
	}
	return out.Pop().(string)
}

func main() {
	push("1")
	push("2")
	fmt.Println(pop())
	push("3")
	fmt.Println(pop())
	push("4")
	fmt.Println(pop())
	push("5")
	//fmt.Println(pop())
	//fmt.Println(pop())
	//fmt.Println(pop())
}
