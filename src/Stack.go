package main
//
//import (
//	//"fmt"
//	"log"
//)
//
//// 用数组实现stack几乎是性能最优的方式
//type Stack []interface{}
//
//
//func (s Stack) IsEmpty() bool {
//	return len(s) == 0
//}
//
//func (s Stack) Size() int {
//	return len(s)
//}
//
//func (s *Stack) Push(item interface{}) {
//	*s = append(*s, item)
//}
//
//func (s *Stack) Pop() interface{} {
//	n := len(*s) - 1
//	result := (*s)[n]
//	*s = (*s)[:n]
//	return result
//}
//
//func (s *Stack) Top() interface{} {
//	if s.Size() <= 0 {
//		log.Fatal("null stack")
//	}
//	return (*s)[s.Size()-1 ]
//}
//
////func main() {
////	stack := new(Stack)
////	fmt.Println(stack.IsEmpty())
////	//stack.Push("a")
////	//stack.Push("a")
////	//stack.Push("a")
////	//stack.Push("b")
////	//stack.Push("b")
////	//stack.Push("c")
////	//stack.Push("d")
////	//stack.Push("e")
////	fmt.Println(cap(*stack))
////	fmt.Println(stack.Top())
////	//fmt.Println(stack.Pop())
////	//fmt.Println(stack.Pop())
////	//fmt.Println(stack.Pop())
////	//fmt.Println(stack.Pop())
////	//fmt.Println(stack.Pop())
////	//fmt.Println(stack.Pop())
////	//fmt.Println(stack.Size())
////}
