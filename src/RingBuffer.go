package main

import (
	"fmt"
	"sync"
	"time"
)

type RingBuffer struct {
	elements [4]string
	writeOps int //用channel的话 把这两个合并?
	readOps  int
	cond     sync.Cond
}

func (rf *RingBuffer) Put(item string) {
	rf.cond.L.Lock()
	if rf.isFull() {
		rf.cond.Wait()
	}
	rf.elements[rf.writeOps] = item
	fmt.Print("put posi:", rf.writeOps, " ,item ", item)
	rf.writeOps = (rf.writeOps + 1) % len(rf.elements)
	fmt.Println("; now write posi:", rf.writeOps)
	rf.cond.Signal()
	rf.cond.L.Unlock()
}

func (rf *RingBuffer) Take() string {
	rf.cond.L.Lock()
	if rf.isEmpty() {
		rf.cond.Wait()
	}
	result := rf.elements[rf.readOps]
	fmt.Print("take:", result)
	rf.readOps = (rf.readOps + 1) % len(rf.elements)
	fmt.Println("; now read posi:", rf.readOps)
	rf.cond.Signal()
	rf.cond.L.Unlock()
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
	var mutex sync.Mutex
	c := sync.NewCond(&mutex)
	rf := RingBuffer{e, 0, 0, *c}
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
		time.Sleep(5 * time.Second)
		rf.Take()
		rf.Take()
		rf.Take()
		time.Sleep(5 * time.Second)
		rf.Take()
		rf.Take()
		rf.Take()
		time.Sleep(5 * time.Second)
	}()
	time.Sleep(10 * time.Minute)
}
