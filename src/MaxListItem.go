package main

import (
	"fmt"
)

type IntNode struct {
	item int
	next *IntNode
}

func max(first *IntNode) int {
	if first.next == nil {
		return first.item
	}

	max := max(first.next)
	if first.item > max {
		return first.item
	} else {
		return max
	}
}

func main() {
	n1 := IntNode{1, nil}
	n2 := IntNode{4, nil}
	n3 := IntNode{2, nil}
	n4 := IntNode{9, nil}
	n5 := IntNode{0, nil}
	n1.next = &n2
	n2.next = &n3
	n3.next = &n4
	n5.next = nil
	fmt.Println(max(&n1))
}
