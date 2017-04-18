package main

import (
	"fmt"
	"log"
)

type DoubleNode struct {
	item string
	pre  *DoubleNode
	next *DoubleNode
}

var first *DoubleNode
var last *DoubleNode

func insertAtHead(value string) {
	if first == nil {
		first = &DoubleNode{value, nil, nil}
		last = first
		return
	}
	node := &DoubleNode{value, nil, first}
	first.pre = node
	first = node
}

func delAtHead(value string) {
	if first == nil {
		log.Fatal("list nil")
	}
	if first.next == nil {
		first = nil
		last = nil
		return
	}
	node := first.next
	node.pre = nil
	first.next = nil
	first = node
}

func insertAtTail(value string) {
	if last == nil {
		last = &DoubleNode{value, nil, nil}
		first = last
		return
	}
	node := &DoubleNode{value, last, nil}
	last.next = node
	last = node
}

func delAtTail(value string, first *DoubleNode) {
	if last == nil {
		log.Fatal("list nil")
	}
	if last.pre == nil {
		first = nil
		last = nil
		return
	}
	node := last.pre
	node.next = nil
	last.pre = nil
	last = node
}

func insertBefore(posi *DoubleNode, value string) {
	node := DoubleNode{value, posi.pre, posi}
	posi.pre = &node
}

func insertAfter(posi *DoubleNode, value string) {
	node := DoubleNode{value, posi, posi.next}
	posi.next = &node
}

//func delete(value string) {
//
//}
func main() {
	insertAtTail("1")
	insertAtTail("2")
	insertAtTail("3")
	insertAtTail("4")
	insertAtTail("5")
	insertAtTail("6")
	insertAtTail("7")
	insertAtHead("-1")
	insertAtHead("-2")
	insertAtHead("-3")
	fmt.Println(first.item)
	fmt.Println(first.next.item)
	fmt.Println(first.next.next.item)
	fmt.Println(first.next.next.next.item)
	fmt.Println(first.next.next.next.next.item)
	fmt.Println(first.next.next.next.next.next.item)
	fmt.Println(first.next.next.next.next.next.next.pre.pre.item)
	//fmt.Println(first.item)
}
