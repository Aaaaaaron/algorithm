package main

import "net/http"

func reverseIterative(first *Node) *Node {
	reverseFirst := &Node{}
	req http.NewRequest()
	for first != nil {
		second := first.next
		first.next = reverseFirst
		reverseFirst = first
		first = second
	}
	return reverseFirst
}

func reverseRecursive(first *Node) *Node {
	if first == nil {
		return nil
	}
	if first.next == nil {
		return first
	}

	second := first.next
	rest := reverseRecursive(second)
	second.next = first
	first.next = nil
	return rest
}
