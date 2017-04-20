package main

import (
	"fmt"
	"log"
)

type Deque struct {
	item [8]string
	head int
	tail int
}

func (deque *Deque) LPush(value string) { // add firstN
	deque.head = (deque.head - 1) % len(deque.item)
	if deque.head < 0 { // head = -1
		deque.head += len(deque.item)
	}
	if deque.head == deque.tail {
		log.Fatal("full")
	}
	deque.item[deque.head] = value
	fmt.Println("head", deque.head)
}

func (deque *Deque) RPush(value string) { // add last
	if (deque.tail+1)%len(deque.item) == deque.head {
		log.Fatal("full")
	}
	deque.item[deque.tail] = value
	deque.tail = (deque.tail + 1) % len(deque.item)
	fmt.Println("tail", deque.tail)
}

func (deque *Deque) LPop() string {
	head := deque.head
	deque.head = (deque.head + 1) % len(deque.item) //因为不删除元素,所以这里要是pop的比push多,会取到过期数据.还有一种方案是置为""
	return deque.item[head]
}

func (deque *Deque) RPop() string {
	deque.tail = (deque.tail - 1) % len(deque.item)
	if deque.tail < 0 {
		deque.tail += len(deque.item)
	}
	return deque.item[deque.tail]
}

func main() {
	var item [8]string
	deque := Deque{item, 0, 0}
	deque.LPush("1")
	deque.LPush("2")
	deque.LPush("3")
	deque.RPush("4")
	deque.RPush("5")
	deque.RPush("6")
	deque.RPush("7")
	fmt.Println(deque.RPop())
	fmt.Println(deque.RPop())
	fmt.Println(deque.RPop())
	fmt.Println(deque.RPop())
	fmt.Println(deque.RPop())
	fmt.Println(deque.RPop())
	//fmt.Println(-1&7)
	//deque.LPush("4")
}
