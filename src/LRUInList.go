package main

import "fmt"

type Node struct {
	item string
	next *Node
}

var firstN *Node

func put(item string) {
	node := firstN
	for node != nil {
		if item == node.item {
			delItem(node)
			break
		}
		node = node.next
	}
	firstN = &Node{item, firstN}
}

func get(item string) bool {
	isFound := false
	node := firstN
	for node != nil {
		if item == node.item {
			delItem(node)
			firstN = &Node{item, firstN}
			isFound = true
			break
		}
		node = node.next
	}
	return isFound
}

func delItem(node *Node) {
	if node.next != nil {
		nextNode := node.next
		node.item = nextNode.item
		node.next = nextNode.next
		nextNode.next = nil
	}
}

func print() {
	node := firstN
	for node != nil {
		fmt.Println(node.item)
		node = node.next
	}
}

func main() {
	firstN = &Node{"1", nil}
	put("a")
	put("b")
	put("c")
	put("d")
	//put("a")
	get("a")
	get("d")
	//a := &Node{"a", nil}
	//b := &Node{"b", nil}
	//c := &Node{"c", nil}
	//firstN.next = a
	//a.next = b
	//b.next = c
	print()
	//fmt.Println(firstN.item)
}
