package main

import (
	"fmt"
	"time"
)

type RingBuffer struct {
	elements  [4]string
	writeOps  int //用channel的话 把这两个合并?
	readOps   int
	capacityC chan int
}

func (rf *RingBuffer) Put(item string) {
	rf.capacityC <- 1
	rf.elements[rf.writeOps] = item
	fmt.Print("put posi:", rf.writeOps, " ,item ", item)
	rf.writeOps = (rf.writeOps + 1) % len(rf.elements)
	fmt.Println("; now write posi:", rf.writeOps)
}

func (rf *RingBuffer) Take() string {
	<-rf.capacityC
	result := rf.elements[rf.readOps]
	fmt.Print("take:", result)
	rf.readOps = (rf.readOps + 1) % len(rf.elements)
	fmt.Println("; now read posi:", rf.readOps)
	return result
}

func (rf *RingBuffer) isFull() bool { // mutex style
	return (rf.writeOps+1)%len(rf.elements) == rf.readOps
}

func (rf *RingBuffer) isEmpty() bool { // mutex style
	return (rf.readOps+1)%len(rf.elements) == rf.writeOps
}

func main() {
	var e [4]string
	c := make(chan int, 3)
	rf := RingBuffer{e, 0, 0, c}
	go func() {
		rf.Put("1")
		rf.Put("2")
		rf.Put("3")
		rf.Put("4")
		rf.Put("5")
		rf.Put("6")
		rf.Put("7")
		rf.Put("8")
		rf.Put("9")
		rf.Put("10")
	}()
	go func() {
		time.Sleep(2 * time.Second)
		rf.Take()
		rf.Take()
		rf.Take()
		time.Sleep(2 * time.Second)
		rf.Take()
		rf.Take()
		rf.Take()
		time.Sleep(2 * time.Second)
		rf.Take()
		rf.Take()
		rf.Take()
		time.Sleep(2 * time.Second)
	}()
	time.Sleep(10 * time.Minute)
}
