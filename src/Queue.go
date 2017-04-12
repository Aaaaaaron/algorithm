package main

import (
	"fmt"
	"log"
)

// 用链表实现队列几乎是最优的方式

type Node struct {
	item string
	next *Node
}

type Queue struct {
	first *Node
	last  *Node
	size  int
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}

func (q *Queue) enqueue(msg string) {
	newNode := Node{msg, nil}
	if q.last != nil {
		q.last.next = &newNode
		q.last = &newNode
	} else {
		q.last = &Node{msg, nil}
		q.first = q.last
	}

	q.size++
}

func (q *Queue) dequeue() string {
	if q.first == nil {
		log.Fatal("nil")
	}
	r := q.first.item
	q.first = q.first.next
	q.size--
	return r
}

func main() {
	var q Queue
	q.enqueue("a")
	q.enqueue("b")
	q.enqueue("c")
	q.enqueue("d")
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
	fmt.Println(q.dequeue())
}
